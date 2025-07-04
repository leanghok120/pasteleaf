package server

import (
	"net/http"

	"github.com/leanghok120/pasteleaf/internal/handlers"
)

func New() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("GET /pastes", handlers.GetPastes)
	mux.HandleFunc("GET /pastes/{id}", handlers.GetPaste)
	mux.HandleFunc("POST /pastes", handlers.CreatePaste)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
