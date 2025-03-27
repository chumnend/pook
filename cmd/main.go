package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chumnend/pook/internal/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", utils.Pong)

	mux.HandleFunc("POST /v1/register", utils.NotImplemented)
	mux.HandleFunc("POST /v1/login", utils.NotImplemented)
	mux.HandleFunc("GET /v1/users/{user_id}", utils.NotImplemented)

	mux.HandleFunc("POST /v1/books", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}", utils.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}", utils.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}", utils.NotImplemented)

	mux.HandleFunc("POST /v1/books/{book_id}/pages", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/pages", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/pages/{page_id}", utils.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}/pages/{page_id}", utils.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}/pages/{page_id}", utils.NotImplemented)

	mux.HandleFunc("POST /v1/books/{book_id}/comments", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/comments", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/comments/{comment_id}", utils.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}/comments/{comment_id}", utils.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}/comments/{comment_id}", utils.NotImplemented)

	mux.HandleFunc("POST /v1/books/{book_id}/ratings/", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/ratings/", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/ratings/{rating_id}", utils.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}/ratings/{rating_id}", utils.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}/ratings/{rating_id}", utils.NotImplemented)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on port", port)
	http.ListenAndServe(":"+port, mux)
}
