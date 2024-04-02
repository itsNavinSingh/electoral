package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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
func (m *Repository) PartyDonor(w http.ResponseWriter, r *http.Request){
	var Datavar model.DonorPartyData
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	Datavar.Title = r.FormValue("name")
	party := strings.ReplaceAll(Datavar.Title, "'", "''")
	query := fmt.Sprintf("SELECT name_purchaser, SUM(denominations) AS total_denominations FROM matched WHERE name_party = '%s' GROUP BY name_purchaser ORDER BY total_denominations DESC;", party)
	var rows *sql.Rows
	rows,  err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var partydata model.TotalWise
		partydata.Title = Datavar.Title
		err = rows.Scan(&partydata.Name, &partydata.Amount)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
		Datavar.Data = append(Datavar.Data, partydata)
	}
	render.RenderTemplate(w, r, "partydonor.page.tmpl", Datavar)
}
func (m *Repository) PartyDonorDate(w http.ResponseWriter, r *http.Request){
	var dotas model.PartyDonorDateDatas
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	dotas.Party = r.FormValue("party")
	dotas.Donor = r.FormValue("donor")
	party := strings.ReplaceAll(dotas.Party, "'", "''")
	donor := strings.ReplaceAll(dotas.Donor, "'", "''")
	query := fmt.Sprintf("SELECT to_char(date_encashment, 'DD-MM-YYYY'), SUM(denominations) AS total_denominations FROM matched WHERE name_party='%s' AND name_purchaser='%s' GROUP BY date_encashment ORDER BY date_encashment ASC;", party, donor)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Something Went wrong", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var details model.PartyDonorDate
		err = rows.Scan(&details.Date, &details.Amount)
		details.Party = dotas.Party
		details.Donor = dotas.Donor
		if err != nil {
			http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
			return
		}
		dotas.Data = append(dotas.Data, details)
	}
	render.RenderTemplate(w, r, "partydonordate.page.tmpl", dotas)
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
func (m *Repository) DonorParty(w http.ResponseWriter, r *http.Request){
	var Datavar model.DonorPartyData
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	Datavar.Title = r.FormValue("donor")
	donor := strings.ReplaceAll(Datavar.Title, "'", "''")
	query := fmt.Sprintf("SELECT name_party, SUM(denominations) AS total_denominations FROM matched WHERE name_purchaser = '%s' GROUP BY name_party ORDER BY total_denominations DESC;", donor)
	var rows *sql.Rows
	rows,  err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var partydata model.TotalWise
		partydata.Title = Datavar.Title
		err = rows.Scan(&partydata.Name, &partydata.Amount)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
		Datavar.Data = append(Datavar.Data, partydata)
	}
	render.RenderTemplate(w, r, "donorparty.page.tmpl", Datavar)
}
func (m *Repository) DonorPartyDate(w http.ResponseWriter, r *http.Request){
	var dotas model.PartyDonorDateDatas
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	dotas.Party = r.FormValue("party")
	dotas.Donor = r.FormValue("donor")
	party := strings.ReplaceAll(dotas.Party, "'", "''")
	donor := strings.ReplaceAll(dotas.Donor, "'", "''")
	query := fmt.Sprintf("SELECT to_char(date_purchase, 'DD-MM-YYYY'), SUM(denominations) AS total_denominations FROM matched WHERE name_party='%s' AND name_purchaser='%s' GROUP BY date_purchase ORDER BY date_purchase ASC;", party, donor)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Something Went wrong", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var details model.PartyDonorDate
		err = rows.Scan(&details.Date, &details.Amount)
		details.Party = dotas.Party
		details.Donor = dotas.Donor
		if err != nil {
			http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
			return
		}
		dotas.Data = append(dotas.Data, details)
	}
	render.RenderTemplate(w, r, "donorpartydate.page.tmpl", dotas)
}
func (m *Repository) DonorDetails(w http.ResponseWriter, r *http.Request) {
	var DonorDatas model.DonorDetails
	err := r.ParseForm()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	// to do remove double double quote to single quote
	donor := r.FormValue("donor")
	DonorDatas.Title = strings.ReplaceAll(donor, "'", "''")
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
	DonorDatas.Title = donor
	render.RenderTemplate(w, r, "donordetails.page.tmpl", DonorDatas)
}
func (m *Repository) Details(w http.ResponseWriter, r *http.Request) {
	var BondDetails model.Matched
	// to do
	// problem in fetching the bond no
	err := r.ParseForm()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	bondno := r.FormValue("BondNo")
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
