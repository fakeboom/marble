/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"reflect"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/securekey/marbles-perf/api"
	"github.com/securekey/marbles-perf/fabric-client"
	"github.com/securekey/marbles-perf/utils"
)


func getExpert(w http.ResponseWriter, r *http.Request) {
	var expert api.Expert
	getEntity(w, r, &expert)
}
func getInstitution(w http.ResponseWriter, r *http.Request) {
	var institution api.Institution
	getEntity(w, r, &institution)
}
func getCity(w http.ResponseWriter, r *http.Request) {
	var city api.City
	getEntity(w, r, &city)
}
func getDemand(w http.ResponseWriter, r *http.Request) {
	var demand api.Demand
	getEntity(w, r, &demand)
}
func getScheme(w http.ResponseWriter, r *http.Request) {
	var scheme api.Scheme
	getEntity(w, r, &scheme)
}
func getPatent(w http.ResponseWriter, r *http.Request) {
	var patent api.Patent
	getEntity(w, r, &patent)
}
func getPaper(w http.ResponseWriter, r *http.Request) {
	var paper api.Paper
	getEntity(w, r, &paper)
}



func change(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here1")
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "failed to read request body: %s", err)
		return
	}
    type Typecheck struct{
		Type string `json:"type"`
	}
	var typec Typecheck
	if err := json.Unmarshal(payload, &typec); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "failed to parse payload json2: %s", err)
		return
	}
	type Idcheck struct{
		Id string `json:"id"`
	}
	var  idcheck Idcheck
	if err := json.Unmarshal(payload, &idcheck); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "failed to parse payload json0: %s", err)
		return
	}
	fmt.Println("here2")
	var owner interface{}
	switch typec.Type{
		case "expert" : owner = &api.Expert{}
		case "institution" : owner = &api.Institution{}
		case "city" : owner = &api.City{}
		case "demand" : owner = &api.Demand{}
		case "scheme" : owner = &api.Scheme{}
		case "patent" : owner = &api.Patent{}
		case "paper"  : owner = &api.Paper{}
		case "transfer"  : owner = &api.Transfer{}
		
	}

	if err := json.Unmarshal(payload, owner); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "failed to parse payload json1: %s", err)
		return
	}
	fmt.Println("here3")
	response, err := dochange(owner, typec.Type, idcheck.Id)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSONResponse(w, http.StatusOK, response)
}
func dochange(marble interface{} , Type string, Id string) (resp api.Response, err error) {
	fmt.Println("here4"+Type)
	id := Id
	if id == "" {
		id, err = utils.GenerateRandomAlphaNumericString(31)
		if err != nil {
			err = fmt.Errorf("failed to generate random string for id: %s", err)
			return
		}
	}
	if Type == "patent" {id = "P" +id;
	}else { 
		id = string(Type[0]) + id;
	}
	fmt.Println("here5"+id)
	args := []string{
		"change",
		id,
	}
	rVal := reflect.ValueOf(marble).Elem()
	for  i := 1 ;i< rVal.NumField(); i++{
		args = append(args, rVal.Field(i).String())
	}


	data, ccErr := fc.InvokeCC(ConsortiumChannelID, MarblesCC, args, nil)
	if ccErr != nil {
		err = fmt.Errorf("cc invoke failed: %s: %v", err, args)
		return
	}

	resp = api.Response{
		Id:   id,
		TxId: data.FabricTxnID,
	}
	return
}


// transfer transfers marble ownership
//
func transfer(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "failed to read request body: %s", err)
		return
	}

	var transfer api.Transfer
	if err := json.Unmarshal(payload, &transfer); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "failed to parse payload json: %s", err)
		return
	}

	args := []string{
		"set_owner",
		transfer.MarbleId,
		transfer.ToOwnerId,
	}

	data, err := fc.InvokeCC(ConsortiumChannelID, MarblesCC, args, nil)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "cc invoke failed: %s: %v", err, args)
		return
	}
	response := api.Response{
		Id:   transfer.MarbleId,
		TxId: data.FabricTxnID,
	}
	writeJSONResponse(w, http.StatusOK, response)
}

// clearMarbles remove all marbles from ledger
//
func clearMarbles(w http.ResponseWriter, r *http.Request) {
	response, err := doClearMarbles()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	writeJSONResponse(w, http.StatusOK, response)
}

func doClearMarbles() (response api.ClearMarblesResponse, err error) {
	args := []string{"clear_marbles"}
	data, ccErr := fc.InvokeCC(ConsortiumChannelID, MarblesCC, args, nil)
	if ccErr != nil {
		err = fmt.Errorf("cc invoke failed: %s: %v", ccErr, args)
		return
	}

	if err = json.Unmarshal(data.Payload, &response); err != nil {
		err = fmt.Errorf("failed to JSON unmarshal cc response: %s: %v: %s", err, args, data.Payload)
		return
	}
	response.TxId = data.FabricTxnID
	return
}

// getEntity retrieves an existing entity
//
func getEntity(w http.ResponseWriter, r *http.Request, entity interface{}) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		writeErrorResponse(w, http.StatusBadRequest, "id not provided")
		return
	}

	data, err := doGetEntity(id, entity)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(data) == 0 {
		writeErrorResponse(w, http.StatusNotFound, "id not found")
		return
	}

	writeJSONResponse(w, http.StatusOK, entity)
}

func doGetEntity(id string, entity interface{}) ([]byte, error) {
	args := []string{
		"read",
		id,
	}

	data, err := fc.QueryCC(0, ConsortiumChannelID, MarblesCC, args, nil)
	if err != nil {
		return nil, fmt.Errorf("cc invoke failed: %s", err)
	}

	payloadJSON := data.Payload

	if len(payloadJSON) > 0 && entity != nil {
		if err := json.Unmarshal([]byte(payloadJSON), entity); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cc response payload: %s: %s", err, payloadJSON)
		}
	}
	return payloadJSON, nil
}

func doGetOwner(id string) (*api.Owner, error) {
	var owner api.Owner
	if data, err := doGetEntity(id, &owner); err != nil {
		return nil, err
	} else if len(data) == 0 {
		return nil, nil
	}
	return &owner, nil
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		writeErrorResponse(w, http.StatusBadRequest, "id not provided")
		return
	}

	response, err := dodelete(id)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSONResponse(w, http.StatusOK, response)
}
func dodelete(id string) (resp api.Response, err error) {
	args := []string{
		"delete_marble",
		id,
	}

	data, ccErr := fc.InvokeCC(ConsortiumChannelID, MarblesCC, args, nil)
	if ccErr != nil {
		err = fmt.Errorf("cc invoke failed: %s: %v", err, args)
		return
	}

	resp = api.Response{
		Id:   id,
		TxId: data.FabricTxnID,
	}
	return
}
func read_everything(w http.ResponseWriter, r *http.Request){
	type Everything struct {
		Owners  		[]api.Owner  		`json:"owners"`
		Marbles 		[]api.Marble 		`json:"marbles"`
		Experts			[]api.Expert 		`json:"experts"`
		Institutions 	[]api.Institution	`json:"institutions"`
		Citys			[]api.City			`json:"citys"`
		Demands			[]api.Demand		`json:"demands"`
		Schemes			[]api.Scheme		`json:"schemes"`
		Patents			[]api.Patent		`json:"patents"`
		Papers			[]api.Paper			`json:"papers"`
	}
	args := []string{
		"read_everything",
	}

	data, ccErr := fc.InvokeCC(ConsortiumChannelID, MarblesCC, args, nil)
	if ccErr != nil {
		fmt.Errorf("cc invoke failed: %s: %v", ccErr, args)
		return
	}
	var er Everything
	err := json.Unmarshal(data.Payload,&er)
	if err != nil {
		fmt.Errorf("Unmarshal error in everything", err)
		return
	}
	writeJSONResponse(w, http.StatusOK, er)
}
func get_history(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		writeErrorResponse(w, http.StatusBadRequest, "id not provided")
		return
	}
	args := []string{
		"getHistory",
		id,
	}

	data, ccErr := fc.InvokeCC(ConsortiumChannelID, MarblesCC, args, nil)
	if ccErr != nil {
		fmt.Errorf("cc invoke failed: %s: %v", ccErr, args)
		return
	}
	var er []api.AuditHistory
	err := json.Unmarshal(data.Payload,&er)
	if err != nil {
		fmt.Errorf("Unmarshal error in everything", err)
		return
	}
	writeJSONResponse(w, http.StatusOK, er)
}