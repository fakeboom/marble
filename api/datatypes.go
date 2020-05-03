/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package api

// Marble data structure
//
type Marble struct {
	Id             string `json:"id"` //the fieldtags are needed to keep case from bouncing around
	Color          string `json:"color"`
	Size           int    `json:"size"` //size in mm of marble
	Owner          Owner  `json:"owner"`
	AdditionalData string `json:"additionalData,omitempty"`
}

// Owner (user) of a marble
//
type Owner struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Company  string `json:"company"`
}

// Transfer reprensents an ownership transfer request
//
type Transfer struct {
	MarbleId    string `json:"marbleId"`
	ToOwnerId   string `json:"toOwnerId"`
	AuthCompany string `json:"authCompany"` // should be fromOwner's company
}

type Expert struct { //专家
	Id    			string 		`json:"id"`
	ExpertID    	string   	`json:"expertid"`  
	ExpertName 		string  	`json:"expername"` 
	Introduction 	string   	`json:"introduction"`
	Affiliation  	string   	`json:"affiliation"`
	E-mail   		string		`json:"e-mail"`
	Telephone 		string 		`json:"telephone"`
	Fax				string 		`json:"fax"` 
	Pwd				string		`json:"pwd"` 
}
type Institution struct{//单位
	Id	            string 		`json:"id"`
	InstitutionID	string		`json:"institutionid"`
	InstitutionName	string		`json:"institutionname"`
	Introduction	string		`json:"introdution"`
	Address			string		`json:"address"`
	E-mail   		string		`json:"e-mail"`
	Telephone 		string 		`json:"telephone"`
	Fax				string 		`json:"fax"` 
	Pwd				string		`json:"pwd"` 

}
type City  struct{//城市
	Id				string		`json:"id"`
	CityID			string		`json:"cityid"`
	CityName		string		`json:"cityname"`
	CityLevel		string		`json:"citylevel"`
	NetworkLink		string		`json:"networklink"`
	E-mail   		string		`json:"e-mail"`
	Telephone 		string 		`json:"telephone"`
	Fax				string 		`json:"fax"` 
	Pwd				string		`json:"pwd"` 
}
type Demand struct{//项目需求
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
}
type Scheme	struct{//解决方案
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

}
type Patent struct{//专利
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
}
type Paper struct{//论文
	Id				string		`json:"id"`
	OwnerId			string		`json:"ownerid"`
	PaperID			string		`json:"paperid"`
	PaperTitle		string		`json:"papertitle"`
	PAbstract		string		`json:"padstract"`
	PKeyword		string		`json:"pkeyword"`
	PDate			string		`json:"pdate"`
	PFile			string		`json:"pfile"`
	DomainID		string		`json:"domainid"`
}

// Response data structure for entity creation or transfer
//
type Response struct {
	Id    string `json:"id"`              // entity id (owner or marble)
	TxId  string `json:"txId"`            // fabric transaction id
	Error string `json:"error,omitempty"` // error message if any from chaincode
}

type ClearMarblesResponse struct {
	TxId    string `json:"txId"`            // fabric transaction id
	Error   string `json:"error,omitempty"` // error message if any from chaincode
	Found   int    `json:"found"`
	Deleted int    `json:"deleted"`
}

type InitBatchRequest struct {
	Concurrency     int  `json:"concurrency"`     // concurrency
	Iterations      int  `json:"iterations"`      //# iterations indicates the number of marbles transfers that each worker performs
	DelaySeconds    int  `json:"delaySeconds"`    // delay_seconds indicates the time the worker will wait between transfers
	ClearMarbles    bool `json:"clearMarbles"`    // clearMarbles indicates whether the client will delete all marbles from the ledger prior to the test
	ExtraDataLength int  `json:"extraDataLength"` // extraDataLength specifies the size of extra data attached to the marble to increase block size
}

type InitBatchResponse struct {
	BatchID string `json:"batchId"`
}

type BatchResult struct {
	Request                InitBatchRequest `json:"request"`
	Status                 string           `json:"status"`
	TotalSuccesses         int              `json:"totalSuccesses"`
	TotalFailures          int              `json:"totalFailures"`
	TotalSuccessSeconds    int              `json:"totalSuccessSeconds"`
	AverageTransferSeconds float64          `json:"averageTransferSeconds"`
	MinTransferSeconds     float64          `json:"minTransferSeconds"`
	MaxTransferSeconds     float64          `json:"maxTransferSeconds"`
}
