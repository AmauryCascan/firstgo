package handlers

import (
	"html/template"
	"net/http"
)

// route home
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

// route about
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

// fonction qui gère la direction de chage page
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //erreur 500
		return
	}
	t.Execute(w, nil)
}
