package server

import (
	"net/http"

	"github.com/leanghok120/pasteleaf/internal/handlers"
)

func New() *http.Server {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))

	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("GET /pastes", handlers.GetPastes)
	mux.HandleFunc("GET /pastes/{id}", handlers.GetPaste)
	mux.HandleFunc("POST /pastes", handlers.CreatePaste)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
