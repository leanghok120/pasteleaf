package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/leanghok120/pasteleaf/internal/models"
)

func CreatePaste(w http.ResponseWriter, r *http.Request) {
	var paste models.Paste
	if err := json.NewDecoder(r.Body).Decode(&paste); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	id, err := models.GenerateID(9)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paste.ID = id
	paste.CreatedAt = time.Now()
	models.SavePaste(paste)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(paste)
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
