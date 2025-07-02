package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"
)

const idCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateID(length int) (string, error) {
	id := make([]byte, length)
	for i := range id {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(idCharset))))
		if err != nil {
			return "", err
		}
		id[i] = idCharset[n.Int64()]
	}
	return string(id), nil
}

type Paste struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var pastes = make(map[string]Paste)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /pastes", createPost)

	fmt.Println("server listening on :8080")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to pasteleaf!")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var paste Paste
	err := json.NewDecoder(r.Body).Decode(&paste)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := generateID(9)
	if err != nil {
		http.Error(w, "failed to generate id", http.StatusInternalServerError)
		return
	}
	paste.CreatedAt = time.Now()
	pastes[id] = paste

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data, err := json.Marshal(paste)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
