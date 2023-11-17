package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// Holds all the application-wide configuration
type AppConfig struct {
	UseCache      bool // Use this to ignore cache during development for hot reload
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
