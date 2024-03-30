package model

type Matched struct {
	BondNo         string    `json:"BondNo"`
	URN            string    `json:"URN"`
	DateJournal    string `json:"DateJournal"`
	DatePurchase   string `json:"DatePurchase"`
	DateExpiry     string `json:"DateExpiry"`
	NameDonor      string    `json:"NameDonor"`
	Denomination   int       `json:"Denomination"`
	IssueBranch    int       `json:"IssueBranch"`
	IssueTeller    int       `json:"IssueTeller"`
	Status         string    `json:"Status"`
	DateEncashment string `json:"DateEncashment"`
	NameParty      string    `json:"NameParty"`
	PayBranch      int       `json:"PayBranch"`
	PayTeller      int       `json:"PayTeller"`
}
type TotalWise struct {
	Name   string
	Amount int
}
type DonorDetails struct {
	Title string
	Data  []Matched
}
