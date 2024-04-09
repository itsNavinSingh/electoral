package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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


//Party Related Handler
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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var party model.TotalWise
		err = rows.Scan(&party.Name, &party.Amount)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Parties = append(Parties, party)
	}
	render.RenderTemplate(w, r, "partywise.page.tmpl", Parties)
}

func (m *Repository) PartyDonor(w http.ResponseWriter, r *http.Request){
	var Datavar model.DonorPartyData
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	Datavar.Title = r.FormValue("name")
	party := strings.ReplaceAll(Datavar.Title, "'", "''")
	query := fmt.Sprintf("SELECT name_purchaser, SUM(denominations) AS total_denominations FROM matched WHERE name_party = '%s' GROUP BY name_purchaser ORDER BY total_denominations DESC;", party)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var partydata model.TotalWise
		err = rows.Scan(&partydata.Name, &partydata.Amount)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datavar.Data = append(Datavar.Data, partydata)
	}
	render.RenderTemplate(w, r, "partydonor.page.tmpl", Datavar)
}

func (m *Repository) PartyDonorDate(w http.ResponseWriter, r *http.Request){
	var datas model.PartyDonorDate
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	datas.Party = r.FormValue("party")
	datas.Donor = r.FormValue("donor")
	party := strings.ReplaceAll(datas.Party, "'", "''")
	donor := strings.ReplaceAll(datas.Donor, "'", "''")
	query := fmt.Sprintf("SELECT to_char(date_encashment, 'DD-MM-YYYY'), SUM(denominations) AS total_denominations FROM matched WHERE name_party='%s' AND name_purchaser='%s' GROUP BY date_encashment ORDER BY date_encashment ASC;", party, donor)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var details model.OnePartyDonorDate
		err = rows.Scan(&details.Date, &details.Amount)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		datas.Data = append(datas.Data, details)
	}
	render.RenderTemplate(w, r, "partydonordate.page.tmpl", datas)
}

func (m *Repository) PartyDonorDateDetails(w http.ResponseWriter, r *http.Request){
	var datas model.PartyDonorDateDetails
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	datas.Party = r.FormValue("party")
	datas.Donor = r.FormValue("donor")
	datas.Date = r.FormValue("date")
	party := strings.ReplaceAll(datas.Party, "'", "''")
	donor := strings.ReplaceAll(datas.Donor, "'", "''")
	date := strings.ReplaceAll(datas.Date, "'", "''")
	query := fmt.Sprintf("SELECT bond_no, urn, to_char(date_purchase, 'DD-MM-YYYY') AS datepurchase, denominations FROM matched WHERE name_purchaser='%s' AND name_party='%s' AND date_encashment = DATE '%s';", donor, party, date)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var details model.OnePartyDonorDateDetails
		err = rows.Scan(&details.BondNo, &details.URN, &details.Date, &details.Denomination)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		datas.Data = append(datas.Data, details)
	}
	render.RenderTemplate(w, r, "partydetails.page.tmpl", datas)
}

//Donor Related Handler
func (m *Repository) Donor(w http.ResponseWriter, r *http.Request) {
	var Donors []model.TotalWise
	query := `
	SELECT name_purchaser, SUM(denominations) AS total_denominations
	FROM donor
	GROUP BY name_purchaser
	ORDER BY total_denominations DESC;
	`
	rows, err := m.App.DataBase.Query(query)
	if err != nil{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var donor model.TotalWise
		err = rows.Scan(&donor.Name, &donor.Amount)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Donors = append(Donors, donor)
	}
	render.RenderTemplate(w, r, "donorwise.page.tmpl", Donors)
}

func (m *Repository) DonorParty(w http.ResponseWriter, r *http.Request){
	var Datavar model.DonorPartyData
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	Datavar.Title = r.FormValue("donor")
	donor := strings.ReplaceAll(Datavar.Title, "'", "''")
	query := fmt.Sprintf("SELECT name_party, SUM(denominations) AS total_denominations FROM matched WHERE name_purchaser = '%s' GROUP BY name_party ORDER BY total_denominations DESC;", donor)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var donordata model.TotalWise
		err = rows.Scan(&donordata.Name, &donordata.Amount)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datavar.Data = append(Datavar.Data, donordata)
	}
	render.RenderTemplate(w, r, "donorparty.page.tmpl", Datavar)
}

func (m *Repository) DonorPartyDate(w http.ResponseWriter, r *http.Request) {
	var datas model.PartyDonorDate
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	datas.Party = r.FormValue("party")
	datas.Donor = r.FormValue("donor")
	party := strings.ReplaceAll(datas.Party, "'", "''")
	donor := strings.ReplaceAll(datas.Donor, "'", "''")
	query := fmt.Sprintf("SELECT to_char(date_purchase, 'DD-MM-YYYY'), SUM(denominations) AS total_denominations FROM matched WHERE name_party='%s' AND name_purchaser='%s' GROUP BY date_purchase ORDER BY date_purchase ASC;", party, donor)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var details model.OnePartyDonorDate
		err = rows.Scan(&details.Date, &details.Amount)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		datas.Data = append(datas.Data, details)
	}
	render.RenderTemplate(w, r, "donorpartydate.page.tmpl", datas)
}

func (m *Repository) DonorPartyDateDetails(w http.ResponseWriter, r *http.Request) {
	var datas model.PartyDonorDateDetails
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	datas.Party = r.FormValue("party")
	datas.Donor = r.FormValue("donor")
	datas.Date = r.FormValue("date")
	party := strings.ReplaceAll(datas.Party, "'", "''")
	donor := strings.ReplaceAll(datas.Donor, "'", "''")
	date := strings.ReplaceAll(datas.Date, "'", "''")
	query := fmt.Sprintf("SELECT bond_no, urn, to_char(date_encashment, 'DD-MM-YYYY') AS datepurchase, denominations FROM matched WHERE name_purchaser='%s' AND name_party='%s' AND date_purchase = DATE '%s';", donor, party, date)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var details model.OnePartyDonorDateDetails
		err = rows.Scan(&details.BondNo, &details.URN, &details.Date, &details.Denomination)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		datas.Data = append(datas.Data, details)
	}
	render.RenderTemplate(w, r, "donordetails.page.tmpl", datas)
}

func (m *Repository) Matched(w http.ResponseWriter, r *http.Request) {
	var Datas model.PartialMatched
	query := `
	SELECT bond_no, name_purchaser, name_party, denominations
	FROM matched
	LIMIT 25;
	`
	rows, err := m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var details model.OnePartialMatched
		err = rows.Scan(&details.BondNo, &details.NameDonor, &details.NameParty, &details.Denomination)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datas.Data = append(Datas.Data, details)
	}
	query = "SELECT name_party FROM matched GROUP BY name_party ORDER BY name_party;"
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var party string
		err = rows.Scan(&party)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datas.Parties = append(Datas.Parties, party)
	}
	query = "SELECT name_purchaser FROM matched GROUP BY name_purchaser ORDER BY name_purchaser;"
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var donor string
		err = rows.Scan(&donor)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datas.Donors = append(Datas.Donors, donor)
	}
	Datas.Offset = len(Datas.Data)
	render.RenderTemplate(w, r, "mached.page.tmpl", Datas)
}

func (m *Repository) POSTMatched(w http.ResponseWriter, r *http.Request) {
	var Datas model.PartialMatched
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var offset int
	offset, err = strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		http.Error(w, "Something Went Wrong", http.StatusBadRequest)
	}
	query := fmt.Sprintf("SELECT bond_no, name_purchaser, name_party, denominations FROM matched OFFSET %d LIMIT 25;", offset)
	rows, err := m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var details model.OnePartialMatched
		err = rows.Scan(&details.BondNo, &details.NameDonor, &details.NameParty, &details.Denomination)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datas.Data = append(Datas.Data, details)
	}
	query = "SELECT name_party FROM matched GROUP BY name_party ORDER BY name_party;"
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var party string
		err = rows.Scan(&party)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datas.Parties = append(Datas.Parties, party)
	}
	query = "SELECT name_purchaser FROM matched GROUP BY name_purchaser ORDER BY name_purchaser;"
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var donor string
		err = rows.Scan(&donor)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datas.Donors = append(Datas.Donors, donor)
	}
	Datas.Offset = offset + len(Datas.Data)
	render.RenderTemplate(w, r, "mached.page.tmpl", Datas)
}

func (m *Repository) GetDetails(w http.ResponseWriter, r *http.Request) {
	var Details model.Matched
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	bondno := strings.ReplaceAll(r.FormValue("id"), "'", "''")
	query := fmt.Sprintf("SELECT bond_no, urn, to_char(date_journal, 'DD-MM-YYYY'), to_char(date_purchase, 'DD-MM-YYYY'), to_char(date_expiry, 'DD-MM-YYYY'), name_purchaser, denominations, issue_branch, issue_teller, status, to_char(date_encashment, 'DD-MM-YYYY'), name_party, pay_branch, pay_teller FROM matched WHERE bond_no = '%s';", bondno)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&Details.BondNo, &Details.URN, &Details.DateJournal, &Details.DatePurchase, &Details.DateExpiry, &Details.NameDonor, &Details.Denomination, &Details.IssueBranch, &Details.IssueTeller, &Details.Status, &Details.DateEncashment, &Details.NameParty, &Details.PayBranch, &Details.PayTeller)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	out, _ := json.Marshal(Details)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	var Details model.SearchReasult
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}
	Details.Heading = strings.ToUpper(r.FormValue("BondNo"))
	bondno := strings.ReplaceAll(Details.Heading, "'", "''")
	query := fmt.Sprintf("SELECT bond_no, urn, to_char(date_journal, 'DD-MM-YYYY'), to_char(date_purchase, 'DD-MM-YYYY'), to_char(date_expiry, 'DD-MM-YYYY'), name_purchaser, denominations, issue_branch, issue_teller, status, to_char(date_encashment, 'DD-MM-YYYY'), name_party, pay_branch, pay_teller FROM matched WHERE bond_no = '%s';", bondno)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var Detail model.Matched
		err = rows.Scan(&Detail.BondNo, &Detail.URN, &Detail.DateJournal, &Detail.DatePurchase, &Detail.DateExpiry, &Detail.NameDonor, &Detail.Denomination, &Detail.IssueBranch, &Detail.IssueTeller, &Detail.Status, &Detail.DateEncashment, &Detail.NameParty, &Detail.PayBranch, &Detail.PayTeller)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Details.Data = append(Details.Data, Detail)
	}
	render.RenderTemplate(w, r, "searchresult.page.tmpl", Details)
}

func (m *Repository) FilterResult(w http.ResponseWriter, r *http.Request){
	var Datas model.PartialMatched
	var queries []string
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}
	bondno := r.FormValue("no")
	if len(bondno) > 0 {
		queries = append(queries, fmt.Sprintf("bond_no = '%s'", strings.ToUpper(strings.ReplaceAll(bondno, "'", "''"))))
	}
	party := r.FormValue("party")
	if len(party) > 0 {
		queries = append(queries, fmt.Sprintf("name_party = '%s'", strings.ToUpper(strings.ReplaceAll(party, "'", "''"))))
	}
	donor := r.FormValue("donor")
	if len(donor) > 0 {
		queries = append(queries, fmt.Sprintf("name_purchaser = '%s'", strings.ToUpper(strings.ReplaceAll(donor, "'", "''"))))
	}
	value := r.FormValue("denominations")
	if len(value) > 0 {
		queries = append(queries, fmt.Sprintf("denominations = %s", strings.ToUpper(strings.ReplaceAll(value, "'", "''"))))
	}
	if len(queries) == 0 {
		http.Redirect(w, r, "/matched-details", http.StatusSeeOther)
		return
	}
	condition := strings.Join(queries, " AND ")
	query := fmt.Sprintf("SELECT bond_no, name_purchaser, name_party, denominations FROM matched WHERE %s;", condition)
	var rows *sql.Rows
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var details model.OnePartialMatched
		err = rows.Scan(&details.BondNo, &details.NameDonor, &details.NameParty, &details.Denomination)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datas.Data = append(Datas.Data, details)
	}
	query = "SELECT name_party FROM matched GROUP BY name_party ORDER BY name_party;"
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var party string
		err = rows.Scan(&party)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datas.Parties = append(Datas.Parties, party)
	}
	query = "SELECT name_purchaser FROM matched GROUP BY name_purchaser ORDER BY name_purchaser;"
	rows, err = m.App.DataBase.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var donor string
		err = rows.Scan(&donor)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Datas.Donors = append(Datas.Donors, donor)
	}
	Datas.Offset = len(Datas.Data)
	render.RenderTemplate(w, r, "filterresult.page.tmpl", Datas)
}