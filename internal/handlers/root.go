package handlers

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🌱 Welcome to Pasteleaf!")
}
