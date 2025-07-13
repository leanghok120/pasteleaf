package handlers

import (
	"net/http"
	"text/template"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/layout.html",
	))
	templ.ExecuteTemplate(w, "layout.html", nil)
}
