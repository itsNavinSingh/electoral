package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/itsNavinSingh/electoral/internal/config"
	"github.com/itsNavinSingh/electoral/internal/handlers"
	"github.com/itsNavinSingh/electoral/internal/render"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	app.Inproduction = false
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.DataBase, err = sql.Open("pgx", "YOUR DATABASE INFO")
	if err != nil {
		log.Fatal(err)
	}
	defer app.DataBase.Close()
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
