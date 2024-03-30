package config

import (
	"database/sql"
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Inproduction  bool
	DataBase      *sql.DB
}
