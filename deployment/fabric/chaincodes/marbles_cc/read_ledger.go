package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================================================================================
// Read - read a generic variable from ledger
//
// Shows Off GetState() - reading a key/value from the ledger
//
// Inputs - Array of strings
//  0
//  key
//  "abc"
//
// Returns - string
// ============================================================================================================================
func read(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, jsonResp string
	var err error
	fmt.Println("starting read")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}

	// input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key) //get the var from ledger
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(jsonResp)
	}

	fmt.Println("- end read")
	return shim.Success(valAsbytes) //send it onward
}

// ============================================================================================================================
// Get everything we need (owners + marbles + companies)
//
// Inputs - none
//
// Returns:
// {
//	"owners": [{
//			"id": "o99999999",
//			"company": "United Marbles"
//			"username": "alice"
//	}],
//	"marbles": [{
//		"id": "m1490898165086",
//		"color": "white",
//		"docType" :"marble",
//		"owner": {
//			"company": "United Marbles"
//			"username": "alice"
//		},
//		"size" : 35
//	}]
// }
// ============================================================================================================================
func read_everything(stub shim.ChaincodeStubInterface) pb.Response {
	type Everything struct {
		Owners  		[]Owner  		`json:"owners"`
		Marbles 		[]Marble 		`json:"marbles"`
		Experts			[]Expert 		`json:"experts"`
		Institutions 	[]Institution	`json:"institutions"`
		Citys			[]City			`json:"citys"`
		Demands			[]Demand		`json:"demands"`
		Schemes			[]Scheme		`json:"schemes"`
		Patents			[]Patent		`json:"patents"`
		Papers			[]Paper			`json:"papers"`
		Transfers       []Transfer 		`json:"transfers"`
	}
	var everything Everything

	// ---- Get All Marbles ---- //
	resultsIterator, err := stub.GetStateByRange("m0", "m9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on marble id - ", queryKeyAsStr)
		var marble Marble
		json.Unmarshal(queryValAsBytes, &marble)                //un stringify it aka JSON.parse()
		everything.Marbles = append(everything.Marbles, marble) //add this marble to the list
	}
	fmt.Println("marble array - ", everything.Marbles)

	// ---- Get All Owners ---- //
	ownersIterator, err := stub.GetStateByRange("o0", "o9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer ownersIterator.Close()

	for ownersIterator.HasNext() {
		aKeyValue, err := ownersIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on owner id - ", queryKeyAsStr)
		var owner Owner
		json.Unmarshal(queryValAsBytes, &owner) //un stringify it aka JSON.parse()

		if owner.Enabled { //only return enabled owners
			everything.Owners = append(everything.Owners, owner) //add this marble to the list
		}
	}
	fmt.Println("owner array - ", everything.Owners)

	//Transfers
	resultsIterator, err = stub.GetStateByRange("t0", "t9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on expert id - ", queryKeyAsStr)
		var marble Transfer
		json.Unmarshal(queryValAsBytes, &marble)                //un stringify it aka JSON.parse()
		everything.Transfers = append(everything.Transfers, marble) //add this marble to the list
	}
	fmt.Println("expert array - ", everything.Transfers)

	//Experts
	resultsIterator, err = stub.GetStateByRange("e0", "e9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on expert id - ", queryKeyAsStr)
		var marble Expert
		json.Unmarshal(queryValAsBytes, &marble)                //un stringify it aka JSON.parse()
		everything.Experts = append(everything.Experts, marble) //add this marble to the list
	}
	fmt.Println("expert array - ", everything.Experts)

	//Citys
	resultsIterator, err = stub.GetStateByRange("c0", "c9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on city id - ", queryKeyAsStr)
		var marble City
		json.Unmarshal(queryValAsBytes, &marble)                //un stringify it aka JSON.parse()
		everything.Citys = append(everything.Citys, marble) //add this marble to the list
	}
	fmt.Println("city array - ", everything.Citys)

	//Institutions
	resultsIterator, err = stub.GetStateByRange("i0", "i9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on institution id - ", queryKeyAsStr)
		var marble Institution
		json.Unmarshal(queryValAsBytes, &marble)                //un stringify it aka JSON.parse()
		everything.Institutions = append(everything.Institutions, marble) //add this marble to the list
	}
	fmt.Println("Institutions array - ", everything.Institutions)

		//demand
	resultsIterator, err = stub.GetStateByRange("d0", "d9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on demand id - ", queryKeyAsStr)
		var marble Demand
		json.Unmarshal(queryValAsBytes, &marble)                //un stringify it aka JSON.parse()
		everything.Demands = append(everything.Demands, marble) //add this marble to the list
	}
	fmt.Println("demand array - ", everything.Demands)

		//scheme
	resultsIterator, err = stub.GetStateByRange("s0", "s9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on Scheme id - ", queryKeyAsStr)
		var marble Scheme
		json.Unmarshal(queryValAsBytes, &marble)                //un stringify it aka JSON.parse()
		everything.Schemes = append(everything.Schemes, marble) //add this marble to the list
	}
	fmt.Println("Scheme array - ", everything.Schemes)

	//Patent
	resultsIterator, err = stub.GetStateByRange("P0", "P9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on Patent id - ", queryKeyAsStr)
		var marble Patent
		json.Unmarshal(queryValAsBytes, &marble)                //un stringify it aka JSON.parse()
		everything.Patents = append(everything.Patents, marble) //add this marble to the list
	}
	fmt.Println("Patent array - ", everything.Patents)

		//Paper
	resultsIterator, err = stub.GetStateByRange("p0", "p9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on Paper id - ", queryKeyAsStr)
		var marble Paper
		json.Unmarshal(queryValAsBytes, &marble)                //un stringify it aka JSON.parse()
		everything.Papers = append(everything.Papers, marble) //add this marble to the list
	}
	fmt.Println("Paper array - ", everything.Papers)

	//change to array of bytes
	everythingAsBytes, _ := json.Marshal(everything) //convert to array of bytes
	return shim.Success(everythingAsBytes)
}

// ============================================================================================================================
// Get history of asset
//
// Shows Off GetHistoryForKey() - reading complete history of a key/value
//
// Inputs - Array of strings
//  0
//  id
//  "m01490985296352SjAyM"
// ============================================================================================================================

func getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	type Oship struct{
		Id  string 	`json:"id"`
		OwnerId string `json:"ownerid"`
	}
	type AuditHistory struct {
		TxId  string `json:"txId"`
		Value Oship `json:"value"`
	}
	var history []AuditHistory
	var marble Oship

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	marbleId := args[0]
	fmt.Printf("- start getHistoryForMarble: %s\n", marbleId)

	// Get History
	resultsIterator, err := stub.GetHistoryForKey(marbleId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		historyData, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		var tx AuditHistory
		tx.TxId = historyData.TxId                 //copy transaction id over
		json.Unmarshal(historyData.Value, &marble) //un stringify it aka JSON.parse()
		if historyData.Value == nil {              //marble has been deleted
			var emptyMarble Oship
			tx.Value = emptyMarble //copy nil marble
		} else {
			json.Unmarshal(historyData.Value, &marble) //un stringify it aka JSON.parse()
			tx.Value = marble                          //copy marble over
		}
		history = append(history, tx) //add this tx to the list
	}
	fmt.Printf("- getHistoryForMarble returning:\n%s", history)

	//change to array of bytes
	historyAsBytes, _ := json.Marshal(history) //convert to array of bytes
	return shim.Success(historyAsBytes)
}

// ============================================================================================================================
// Get history of asset - performs a range query based on the start and end keys provided.
//
// Shows Off GetStateByRange() - reading a multiple key/values from the ledger
//
// Inputs - Array of strings
//       0     ,    1
//   startKey  ,  endKey
//  "marbles1" , "marbles5"
// ============================================================================================================================
func getMarblesByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryResultKey := aKeyValue.Key
		queryResultValue := aKeyValue.Value

		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResultKey)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResultValue))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getMarblesByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
