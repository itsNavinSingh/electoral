package handlers

import (
	"net/http"

	"github.com/itsNavinSingh/electoral/internal/config"
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
	render.RenderTemplate(w, r, "partywise.page.tmpl", nil)
}
func (m *Repository) Donor(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "donorwise.page.tmpl", nil)
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
	render.RenderTemplate(w, r, "donordetails.page.tmpl", nil)
}
func (m *Repository) Details(w http.ResponseWriter, r *http.Request){
	//to do
	// send  a json data
}