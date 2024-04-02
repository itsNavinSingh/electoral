package model

type Matched struct {
	BondNo         string `json:"BondNo"`
	URN            string `json:"URN"`
	DateJournal    string `json:"DateJournal"`
	DatePurchase   string `json:"DatePurchase"`
	DateExpiry     string `json:"DateExpiry"`
	NameDonor      string `json:"NameDonor"`
	Denomination   int    `json:"Denomination"`
	IssueBranch    int    `json:"IssueBranch"`
	IssueTeller    int    `json:"IssueTeller"`
	Status         string `json:"Status"`
	DateEncashment string `json:"DateEncashment"`
	NameParty      string `json:"NameParty"`
	PayBranch      int    `json:"PayBranch"`
	PayTeller      int    `json:"PayTeller"`
}
type TotalWise struct {
	Name   string
	Amount int
	Title  string
}
type DonorPartyData struct {
	Title string
	Data  []TotalWise
}
type DonorDetails struct {
	Title string
	Data  []Matched
}
type PartyDonorDate struct {
	Party string
	Donor string
	Date string
	Amount int
}
type PartyDonorDateDatas struct{
	Party string
	Donor string
	Data []PartyDonorDate
}
