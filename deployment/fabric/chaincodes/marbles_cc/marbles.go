package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Asset Definitions - The ledger will store marbles and owners
// ============================================================================================================================

// ----- Marbles ----- //
type Marble struct {
	ObjectType     string        `json:"docType"` //field for couchdb
	Id             string        `json:"id"`      //the fieldtags are needed to keep case from bouncing around
	Color          string        `json:"color"`
	Size           int           `json:"size"` //size in mm of marble
	Owner          OwnerRelation `json:"owner"`
	AdditionalData string        `json:"additionalData"`
}

// ----- Owners ----- //
type Owner struct {
	ObjectType string `json:"docType"` //field for couchdb
	Id         string `json:"id"`
	Username   string `json:"username"`
	Company    string `json:"company"`
	Enabled    bool   `json:"enabled"` //disabled owners will not be visible to the application
}

type OwnerRelation struct {
	Id       string `json:"id"`
	Username string `json:"username"` //this is mostly cosmetic/handy, the real relation is by Id not Username
	Company  string `json:"company"`  //this is mostly cosmetic/handy, the real relation is by Id not Company
}
type Expert struct { //专家
	ObjectType     string        `json:"docType"`
	Id    			string 		`json:"id"`
	ExpertID    	string   	`json:"expertid"`  
	ExpertName 		string  	`json:"expername"` 
	Introduction 	string   	`json:"introduction"`
	Affiliation  	string   	`json:"affiliation"`
	Email   		string		`json:"email"`
	Telephone 		string 		`json:"telephone"`
	Fax				string 		`json:"fax"` 
	Pwd				string		`json:"pwd"` 
	Able			string		`json:"able"`
}
type Institution struct{//单位
	ObjectType     string        `json:"docType"`
	Id	            string 		`json:"id"`
	InstitutionID	string		`json:"institutionid"`
	InstitutionName	string		`json:"institutionname"`
	Introduction	string		`json:"introdution"`
	Address			string		`json:"address"`
	Email   		string		`json:"email"`
	Telephone 		string 		`json:"telephone"`
	Fax				string 		`json:"fax"` 
	Pwd				string		`json:"pwd"` 
	Able			string		`json:"able"`
}

type City  struct{//城市
	ObjectType     string        `json:"docType"`
	Id				string		`json:"id"`
	CityID			string		`json:"cityid"`
	CityName		string		`json:"cityname"`
	CityLevel		string		`json:"citylevel"`
	NetworkLink		string		`json:"networklink"`
	Email   		string		`json:"email"`
	Telephone 		string 		`json:"telephone"`
	Fax				string 		`json:"fax"` 
	Pwd				string		`json:"pwd"`
	Able			string		`json:"able"`
}
type Demand struct{//项目需求
	ObjectType     string        `json:"docType"`
	Id				string		`json:"id"`
	OwnerId			string		`json:"ownerid"`
	DemandID		string		`json:"demandid"`
	KeyWord			string		`json:"keyword"`
	Budget			string		`json:"budget"`
	AnnouncementTime	string	`json:"announcementtime"`
	TenderTime		string		`json:"tendertime"`
	BidOpeningTime	string		`json:"bidopeningtime"`
	OpeningAddress	string		`json:"openingaddress"`
	ProjectContact	string		`json:"projectcontact"`
	ProjectPhone	string		`json:"projectphone"`
	PurchasingUnit	string		`json:"purchasingunit"`
	PurchasingUnitAdd	string		`json:"purchasingunitadd"`
	PurchasingUnitPhone	string		`json:"purchasingunitphone"`
	Agency			string		`json:"agency"`
	AgencyAdd		string		`json:"agencyadd"`
	AgencyPhone		string		`json:"agencyphone"`
	Resources		string		`json:"resources"`
	Description		string		`json:"description"`
	File			string		`json:"file"`
	Note			string		`json:"note"`
	Able			string		`json:"able"`
}
type Scheme	struct{//解决方案
	ObjectType     string        `json:"docType"`
	Id				string		`json:"id"`
	OwnerId			string		`json:"ownerid"`
	SchemeID		string		`json:"schemeid"`
	SchemeTitle		string		`json:"schemetitle"`
	KeyWord			string		`json:"keyword"`
	Period			string		`json:"period"`
	Supplier		string		`json:"supplier"`
	Budget			string		`json:"budget"`	
	ProjectContact	string		`json:"projectcontact"`
	ProjectPhone	string		`json:"projectphone"`
	Resources		string		`json:"resources"`
	Description		string		`json:"description"`
	File			string		`json:"file"`
	Note			string		`json:"note"`
	Able			string		`json:"able"`
}
type Patent struct{//专利
	ObjectType     string        `json:"docType"`
	Id				string		`json:"id"`
	OwnerId			string		`json:"ownerid"`
	PatentID		string		`json:"patentid"`
	PatentNumber	string		`json:"patentnumber"`
	PType			string		`json:"ptype"`
	PName 			string		`json:"pname"`
	PDate			string		`json:"pdate"`
	POpen 			string		`json:"popen"`
	POpenDate		string		`json:"popendate"`
	PState 			string		`json:"pstate"`
	ApplyID			string		`json:"applyid"`
	DomainID		string		`json:"domainid"`
	Able			string		`json:"able"`
}
type Paper struct{//论文
	ObjectType     string        `json:"docType"`
	Id				string		`json:"id"`
	OwnerId			string		`json:"ownerid"`
	PaperID			string		`json:"paperid"`
	PaperTitle		string		`json:"papertitle"`
	PAbstract		string		`json:"padstract"`
	PKeyword		string		`json:"pkeyword"`
	PDate			string		`json:"pdate"`
	PFile			string		`json:"pfile"`
	DomainID		string		`json:"domainid"`
	Able			string		`json:"able"`
}

type Transfer struct {//交易请求
	ObjectType     string        `json:"docType"`
	Id				string		`json:"id"`
	OwnerId			string		`json:"ownerid"`
	MarbleId    string `json:"marbleId"`
	ToOwnerId   string `json:"toOwnerId"`
	Able        string  `json:"able"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode - %s", err)
	}
}

// ============================================================================================================================
// Init - initialize the chaincode
//
// Marbles does not require initialization, so let's run a simple test instead.
//
// Shows off PutState() and how to pass an input argument to chaincode.
// Shows off GetFunctionAndParameters() and GetStringArgs()
// Shows off GetTxID() to get the transaction ID of the proposal
//
// Inputs - Array of strings
//  ["314"]
//
// Returns - shim.Success or error
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Marbles Is Starting Up")
	funcName, args := stub.GetFunctionAndParameters()
	var number int
	var err error
	txId := stub.GetTxID()

	fmt.Println("Init() is running")
	fmt.Println("Transaction ID:", txId)
	fmt.Println("  GetFunctionAndParameters() function:", funcName)
	fmt.Println("  GetFunctionAndParameters() args count:", len(args))
	fmt.Println("  GetFunctionAndParameters() args found:", args)

	// expecting 1 arg for instantiate or upgrade
	if len(args) == 1 {
		fmt.Println("  GetFunctionAndParameters() arg[0] length", len(args[0]))

		// expecting arg[0] to be length 0 for upgrade
		if len(args[0]) == 0 {
			fmt.Println("  Uh oh, args[0] is empty...")
		} else {
			fmt.Println("  Great news everyone, args[0] is not empty")

			// convert numeric string to integer
			number, err = strconv.Atoi(args[0])
			if err != nil {
				return shim.Error("Expecting a numeric string argument to Init() for instantiate")
			}

			// this is a very simple test. let's write to the ledger and error out on any errors
			// it's handy to read this right away to verify network is healthy if it wrote the correct value
			err = stub.PutState("selftest", []byte(strconv.Itoa(number)))
			if err != nil {
				return shim.Error(err.Error()) //self-test fail
			}
		}
	}

	// showing the alternative argument shim function
	alt := stub.GetStringArgs()
	fmt.Println("  GetStringArgs() args count:", len(alt))
	fmt.Println("  GetStringArgs() args found:", alt)

	// store compatible marbles application version
	err = stub.PutState("marbles_ui", []byte("4.0.1"))
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("Ready for action") //self-test pass
	return shim.Success(nil)
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println(" ")
	fmt.Println("starting invoke, for - " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub)
	} else if function == "read" { //generic read ledger
		return read(stub, args)
	} else if function == "write" { //generic writes to ledger
		return write(stub, args)
	} else if function == "delete_marble" { //deletes a marble from state
		return delete_marble(stub, args)
	} else if function == "init_marble" { //create a new marble
		return init_marble(stub, args)
	} else if function == "set_owner" { //change owner of a marble
		return set_owner(stub, args)
	} else if function == "init_owner" { //create a new marble owner
		return init_owner(stub, args)
	} else if function == "read_everything" { //read everything, (owners + marbles + companies)
		return read_everything(stub)
	} else if function == "getHistory" { //read history of a marble (audit)
		return getHistory(stub, args)
	} else if function == "getMarblesByRange" { //read a bunch of marbles by start and stop id
		return getMarblesByRange(stub, args)
	} else if function == "disable_owner" { //disable a marble owner from appearing on the UI
		return disable_owner(stub, args)
	} else if function == "clear_marbles" { //remove all marbles
		return clear_marbles(stub, args)
	} else if function == "delete_marble_noauth" { //delete a marble without checking auth company
		return delete_marble_noauth(stub, args)
	} else if function == "change" {
		return change(stub, args)
	} else if function == "able" {
		return able_alkind(stub, args)
	}

	// error out
	fmt.Println("Received unknown invoke function name - " + function)
	return shim.Error("Received unknown invoke function name - '" + function + "'")
}

// ============================================================================================================================
// Query - legacy function
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("Unknown supported call - Query()")
}
