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
	//mux.Get("/party/details", handlers.Repo.PartyDetails)

	mux.Get("/donor", handlers.Repo.Donor)
	mux.Post("/donor/party", handlers.Repo.DonorParty)
	mux.Post("/donor/party/date", handlers.Repo.DonorPartyDate)
	//mux.Post("/donor/details", handlers.Repo.DonorDetails)

	mux.Get("/matched-details", handlers.Repo.Matched)
	mux.Post("/matched-details", handlers.Repo.Matched)

	mux.Post("/search", handlers.Repo.Search)
	mux.Post("/get-details", handlers.Repo.Details)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}