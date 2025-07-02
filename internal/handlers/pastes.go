package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/leanghok120/pasteleaf/internal/models"
)

func CreatePaste(w http.ResponseWriter, r *http.Request) {
	var paste models.Paste
	if err := json.NewDecoder(r.Body).Decode(&paste); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	id, err := models.GenerateID(9)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	paste.ID = id
	paste.CreatedAt = time.Now()
	models.SavePaste(paste)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(paste)
}
