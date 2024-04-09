package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/itsNavinSingh/electoral/internal/config"
	"github.com/itsNavinSingh/electoral/internal/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/party", handlers.Repo.Party)
	mux.Post("/party/donor", handlers.Repo.PartyDonor)
	mux.Post("/party/donor/date", handlers.Repo.PartyDonorDate)
	mux.Post("/party/donor/date/details", handlers.Repo.PartyDonorDateDetails)

	mux.Get("/donor", handlers.Repo.Donor)
	mux.Post("/donor/party", handlers.Repo.DonorParty)
	mux.Post("/donor/party/date", handlers.Repo.DonorPartyDate)
	mux.Post("/donor/party/date/details", handlers.Repo.DonorPartyDateDetails)

	mux.Get("/matched-details", handlers.Repo.Matched)
	mux.Post("/matched-details", handlers.Repo.POSTMatched)

	mux.Post("/search", handlers.Repo.Search)
	mux.Post("/get-details", handlers.Repo.GetDetails)

	mux.Post("/filter-result",handlers.Repo.FilterResult)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}