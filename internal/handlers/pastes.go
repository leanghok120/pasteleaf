package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/leanghok120/pasteleaf/internal/models"
)

func CreatePaste(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
	}

	paste := models.Paste{
		Title:     r.FormValue("title"),
		Content:   r.FormValue("content"),
		CreatedAt: time.Now(),
	}
	id, _ := models.GenerateID(9)
	paste.ID = id
	models.SavePaste(paste)
}

func GetPaste(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	paste, ok := models.GetPaste(id)
	if !ok {
		http.Error(w, "paste doesn't exist", http.StatusNotFound)
		return
	}

	templ := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/paste.html",
	))
	w.Header().Set("Content-Type", "text/html")
	templ.ExecuteTemplate(w, "layout.html", paste)
}

func GetPastes(w http.ResponseWriter, r *http.Request) {
	pastes := models.GetPastes()

	if len(pastes) == 0 {
		fmt.Fprintln(w, "no pastes available")
		return
	}

	templ := template.Must(template.ParseFiles("templates/pastes.html"))
	w.Header().Set("Content-Type", "text/html")
	templ.Execute(w, pastes)
}
