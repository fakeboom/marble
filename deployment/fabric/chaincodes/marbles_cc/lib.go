package main

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ============================================================================================================================
// Get Marble - get a marble asset from ledger
// ============================================================================================================================
func get_marble(stub shim.ChaincodeStubInterface, id string) (Marble, error) {
	var marble Marble
	marbleAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
	if err != nil {                         //this seems to always succeed, even if key didn't exist
		return marble, errors.New("Failed to find marble - " + id)
	}
	json.Unmarshal(marbleAsBytes, &marble) //un stringify it aka JSON.parse()

	if marble.Id != id { //test if marble is actually here or just nil
		return marble, errors.New("Marble does not exist - " + id)
	}

	return marble, nil
}

// ============================================================================================================================
// Get Owner - get the owner asset from ledger
// ============================================================================================================================
func get_owner(stub shim.ChaincodeStubInterface, id string) (Owner, error) {
	var owner Owner
	ownerAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
		if err != nil {                        //this seems to always succeed, even if key didn't exist
			return owner, errors.New("Failed to get owner - " + id)
		}
		json.Unmarshal(ownerAsBytes, &owner) //un stringify it aka JSON.parse()

	if len(owner.Username) == 0 { //test if owner is actually here or just nil
		return owner, errors.New("Owner does not exist - " + id + ", '" + owner.Username + "' '" + owner.Company + "'")
	}

	return owner, nil
}
//获取专家
func get_expert(stub shim.ChaincodeStubInterface, id string) (Expert, error) {
	var expert Expert;
	ownerAsBytes, err := stub.GetState(id) 
	if err != nil {                        
		return expert, errors.New("Failed to get expert - " + id)
	}
	json.Unmarshal(ownerAsBytes, &expert) 
	if len(expert.ExpertID) == 0 { 
		return expert, errors.New("Expert does not exist - " + id + ", '" + expert.ExpertID )
	}

	return expert, nil
}
//获取单位
func get_institution(stub shim.ChaincodeStubInterface, id string) (Institution, error) {
	var institution Institution;
	ownerAsBytes, err := stub.GetState(id) 
	if err != nil {                        
		return institution, errors.New("Failed to get institution - " + id)
	}
	json.Unmarshal(ownerAsBytes, &institution) 
	if len(institution.InstitutionID) == 0 { 
		return institution, errors.New("institution does not exist - " + id + ", '" + institution.InstitutionID )
	}

	return institution, nil
}
//获取城市
func get_city(stub shim.ChaincodeStubInterface, id string) (City, error) {
	var city City;
	ownerAsBytes, err := stub.GetState(id) 
	if err != nil {                        
		return city, errors.New("Failed to get city - " + id)
	}
	json.Unmarshal(ownerAsBytes, &city) 
	if len(city.CityID) == 0 { 
		return city, errors.New("city does not exist - " + id + ", '" + city.CityID )
	}
	return city, nil
}

//获取需求
func get_demand(stub shim.ChaincodeStubInterface, id string) (Demand, error) {
	var demand Demand;
	ownerAsBytes, err := stub.GetState(id) 
	if err != nil {                        
		return demand, errors.New("Failed to get demand - " + id)
	}
	json.Unmarshal(ownerAsBytes, &demand) 
	if len(demand.DemandID) == 0 { 
		return demand, errors.New("demand does not exist - " + id + ", '" + demand.DemandID)
	}
	return demand, nil
}

//获取解决方案
func get_scheme(stub shim.ChaincodeStubInterface, id string) (Scheme, error) {
	var scheme Scheme;
	ownerAsBytes, err := stub.GetState(id) 
	if err != nil {                        
		return scheme, errors.New("Failed to get scheme - " + id)
	}
	json.Unmarshal(ownerAsBytes, &scheme) 
	if len(scheme.SchemeID) == 0 { 
		return scheme, errors.New("scheme does not exist - " + id + ", '" + scheme.SchemeID)
	}
	return scheme, nil
}

//获取专利
func get_patent(stub shim.ChaincodeStubInterface, id string) (Patent, error) {
	var patent Patent;
	ownerAsBytes, err := stub.GetState(id) 
	if err != nil {                        
		return patent, errors.New("Failed to get patent - " + id)
	}
	json.Unmarshal(ownerAsBytes, &patent) 
	if len(patent.PatentID) == 0 { 
		return patent, errors.New("patent does not exist - " + id + ", '" + patent.PatentID)
	}
	return patent, nil
}
//获取论文
func get_paper(stub shim.ChaincodeStubInterface, id string) (Paper, error) {
	var paper Paper;
	ownerAsBytes, err := stub.GetState(id) 
	if err != nil {                        
		return paper, errors.New("Failed to get paper - " + id)
	}
	json.Unmarshal(ownerAsBytes, &paper) 
	if len(paper.PaperID) == 0 { 
		return paper, errors.New("paper does not exist - " + id + ", '" + paper.PaperID)
	}
	return paper, nil
}
// 检测输入是否为空
func sanitize_arguments(strs []string) error {
	for i, val := range strs {
		if len(val) <= 0 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be a non-empty string")
		}
	}
	return nil
}
