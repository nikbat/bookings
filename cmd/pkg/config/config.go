package config

import (
	"html/template"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	InProduction   bool
	SameSite       http.SameSite
	SessionManager *scs.SessionManager
}
