package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/itsNavinSingh/electoral/internal/config"
	"github.com/itsNavinSingh/electoral/internal/model"
	"github.com/itsNavinSingh/electoral/internal/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", nil)
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.tmpl", nil)
}
func (m *Repository) Party(w http.ResponseWriter, r *http.Request) {
	var Parties []model.TotalWise
	query := `
	SELECT name_party, SUM(denominations) AS total_denominations
	FROM party
	GROUP BY name_party
	ORDER BY total_denominations DESC;
	`
	rows, err := m.App.DataBase.Query(query)
	if err != nil {
		log.Println("Problem in Retriving party name")
	}
	defer rows.Close()
	for rows.Next() {
		var party model.TotalWise
		err = rows.Scan(&party.Name, &party.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		Parties = append(Parties, party)
	}
	render.RenderTemplate(w, r, "partywise.page.tmpl", Parties)
}
func (m *Repository) Donor(w http.ResponseWriter, r *http.Request) {
	var Donors []model.TotalWise
	query := `
	SELECT name_purchaser, SUM(denominations) AS total_denominations
	FROM donor
	GROUP BY name_purchaser
	ORDER BY total_denominations DESC;
	`
	rows, err := m.App.DataBase.Query(query)
	if err != nil {
		log.Println("Problem in retriving donor name")
	}
	defer rows.Close()
	for rows.Next() {
		var donor model.TotalWise
		err = rows.Scan(&donor.Name, &donor.Amount)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		Donors = append(Donors, donor)
	}
	render.RenderTemplate(w, r, "donorwise.page.tmpl", Donors)
}
func (m *Repository) Matched(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "mached.page.tmpl", nil)
}
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "searchresult.page.tmpl", nil)
}
func (m *Repository) PartyDetails(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "partydetails.page.tmpl", nil)
}
func (m *Repository) DonorDetails(w http.ResponseWriter, r *http.Request) {
	var DonorDatas model.DonorDetails
	DonorDatas.Title = r.URL.Query().Get("donor")
	query := fmt.Sprintf("SELECT bond_no, to_char(date_purchase, 'DD-MM-YYYY') AS date_purchase, name_party, denominations FROM matched WHERE name_purchaser = '%s';", DonorDatas.Title)
	rows, err := m.App.DataBase.Query(query)
	if err != nil {
		log.Println("Problem in retriving donor details")
	}
	defer rows.Close()
	for rows.Next() {
		var DonorData model.Matched
		err = rows.Scan(&DonorData.BondNo, &DonorData.DatePurchase, &DonorData.NameParty, &DonorData.Denomination)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		DonorDatas.Data = append(DonorDatas.Data, DonorData)
	}
	render.RenderTemplate(w, r, "donordetails.page.tmpl", DonorDatas)
}
func (m *Repository) Details(w http.ResponseWriter, r *http.Request) {
	bondno := r.URL.Query().Get("BondNo")
	var BondDetails model.Matched
	donorQuery := fmt.Sprintf("SELECT bond_no, urn, to_char(date_journal, 'DD-MM-YYYY') AS date_journal, to_char(date_purchase, 'DD-MM-YYYY') AS date_purchase, to_char(date_expiry, 'DD-MM-YYYY') AS date_expiry, name_purchaser, denominations, issue_branch, issue_teller, status FROM donor WHERE bond_no = '%s';", bondno)
	partyQuery := fmt.Sprintf("SELECT bond_no, to_char(date_encashment, 'DD-MM-YYYY') AS date_encashment, name_party, denominations, pay_branch, pay_teller FROM party WHERE bond_no = '%s';", bondno)
	rows, err := m.App.DataBase.Query(donorQuery)
	if err != nil {
		log.Println("Problem in retriving donor details")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&BondDetails.BondNo, &BondDetails.URN, &BondDetails.DateJournal, &BondDetails.DatePurchase, &BondDetails.DateExpiry, &BondDetails.NameDonor, &BondDetails.Denomination, &BondDetails.IssueBranch, &BondDetails.IssueTeller, &BondDetails.Status)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}

	rows, err = m.App.DataBase.Query(partyQuery)
	if err != nil {
		log.Println("Problem in retriving party details")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&BondDetails.BondNo, &BondDetails.DateEncashment, &BondDetails.NameParty, &BondDetails.Denomination, &BondDetails.PayBranch, &BondDetails.PayTeller)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}

	out, _ := json.Marshal(BondDetails)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
