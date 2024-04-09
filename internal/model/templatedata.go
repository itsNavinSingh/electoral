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
}
type DonorPartyData struct {
	Title string
	Data  []TotalWise
}
type DonorDetails struct {
	Title string
	Data  []Matched
}
type OnePartyDonorDate struct {
	Date string
	Amount int
}
type PartyDonorDate struct{
	Party string
	Donor string
	Data []OnePartyDonorDate
}
type PartyDonorDateDetails struct {
	Party string
	Donor string
	Date string
	Data []OnePartyDonorDateDetails
}
type OnePartyDonorDateDetails struct {
	BondNo string
	URN string
	Date string
	Denomination int
}

type OnePartialMatched struct {
	BondNo string
	NameDonor string
	NameParty string
	Denomination int
}
type PartialMatched struct {
	Parties []string
	Donors  []string
	Offset int
	Data []OnePartialMatched
}
type SearchReasult struct {
	Heading string
	Data []Matched
}