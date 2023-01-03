package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/nikbat/bookings/cmd/pkg/config"
	"github.com/nikbat/bookings/cmd/pkg/models"
)

var app *config.AppConfig

// sets the config for the teamplate package
func SetConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, name string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//below lines are commented because template cache is stored in global config
	//create teamplate cache

	// tc, err := CreateTemplateCache()

	// if err != nil {
	// 	log.Println("Error creating template case ", err)
	// 	return
	// }

	//get requested teamplate
	t, ok := tc[name]
	if !ok {
		log.Fatal("Template %s not found in cache returning ", name)

	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal("Error executing ", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal("Error writing buffer to response ", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	var cachedTemplates = map[string]*template.Template{}

	//get all the templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		log.Println("Error reading all templates from ./templates/*.page.tmpl ", err)
		return cachedTemplates, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Printf("Error parsing teamplate %s at path %s", name, page, err)
			return cachedTemplates, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println("Error reading lauout templates from ./templates/*.page.tmpl ", err)
			return cachedTemplates, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				log.Printf("Error parsing layout teamplat ./templates/*.layout.tmpl", err)
				return cachedTemplates, err
			}
		}

		cachedTemplates[name] = ts

	}

	return cachedTemplates, nil
}

func RenderTemplateOld(w http.ResponseWriter, name string) {
	parsedTemplate, parseError := template.ParseFiles("./templates/"+name, "./templates/base.layout.tmpl")
	if parseError != nil {
		log.Println("Error parsing the template: ", parseError)
	}

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executiong parsed template: ", err)
		return
	}

}
