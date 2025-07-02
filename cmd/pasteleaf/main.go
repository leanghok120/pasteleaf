package main

import (
	"fmt"

	"github.com/leanghok120/pasteleaf/internal/server"
)

func main() {
	server := server.New()
	fmt.Println("server is listening on :8080")
	server.ListenAndServe()
}
