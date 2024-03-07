package handlers

import (
	"bytes"
	"firstgo-project/config"
	"firstgo-project/model"
	"html/template"
	"net/http"
	"path/filepath"
)

// route home
func Home(w http.ResponseWriter, r *http.Request) {
	names := make(map[string]string)
	names["owner"] = "Amaury"

	renderTemplate(w, "home", &models.TemplateData{
		StringData: names,
	})
}

// route about
func About(w http.ResponseWriter, r *http.Request) {
	age := make(map[string]int)
	age["owner"] = 30

	renderTemplate(w, "about", &models.TemplateData{
		IntData: age,
	})
}

var appConfig *config.Config

func CreateTemplates(app *config.Config) {
	appConfig = app
}

// fonction qui gère la direction de chage page
func renderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {

	templateCache := appConfig.TemplateCache

	tmpl, ok := templateCache[tmplName+".page.tmpl"]


	if !ok {
		http.Error(w, "Le template n'existe pas !", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, td)
	buffer.WriteTo(w)

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("../../templates/layouts/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("../../templates/layouts/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	
	return cache, nil
}
