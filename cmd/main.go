package main

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", handlers.Pong)

	mux.HandleFunc("POST /v1/register", handlers.NotImplemented)
	mux.HandleFunc("POST /v1/login", handlers.NotImplemented)
	mux.HandleFunc("GET /v1/users/{user_id}", handlers.NotImplemented)

	mux.HandleFunc("POST /v1/books", handlers.NotImplemented)
	mux.HandleFunc("GET /v1/books", handlers.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}", handlers.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}", handlers.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}", handlers.NotImplemented)

	mux.HandleFunc("POST /v1/books/{book_id}/pages", handlers.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/pages", handlers.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/pages/{page_id}", handlers.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}/pages/{page_id}", handlers.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}/pages/{page_id}", handlers.NotImplemented)

	mux.HandleFunc("POST /v1/books/{book_id}/comments", handlers.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/comments", handlers.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/comments/{comment_id}", handlers.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}/comments/{comment_id}", handlers.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}/comments/{comment_id}", handlers.NotImplemented)

	mux.HandleFunc("POST /v1/books/{book_id}/ratings/", handlers.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/ratings/", handlers.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/ratings/{rating_id}", handlers.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}/ratings/{rating_id}", handlers.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}/ratings/{rating_id}", handlers.NotImplemented)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
