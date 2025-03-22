package main

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/handlers"
)

func main() {
	http.HandleFunc("/ping", handlers.Pong)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
