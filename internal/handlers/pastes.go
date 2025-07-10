package handlers

import (
	"encoding/json"
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

	templ := template.Must(template.ParseFiles("templates/success.html"))
	templ.Execute(w, nil)
}

func GetPaste(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	paste, ok := models.GetPaste(id)
	if !ok {
		http.Error(w, "paste doesn't exist", http.StatusNotFound)
		return
	}

	data, err := json.Marshal(paste)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetPastes(w http.ResponseWriter, r *http.Request) {
	pastes := models.GetPastes()

	if len(pastes) == 0 {
		fmt.Fprintln(w, "no pastes available")
		return
	}

	data, err := json.Marshal(pastes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
