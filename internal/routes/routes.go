package routes

import (
	"net/http"

	"github.com/chumnend/pook/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	// User Routes
	mux.HandleFunc("POST /v1/register", handlers.Register)
	mux.HandleFunc("POST /v1/login", handlers.Login)
	mux.HandleFunc("GET /v1/users/{user_id}", handlers.GetUser)
	// Book Routes
	mux.HandleFunc("POST /v1/books", handlers.CreateBook)
	mux.HandleFunc("GET /v1/books", handlers.GetAllBooks)
	mux.HandleFunc("GET /v1/books/{book_id}", handlers.GetBook)
	mux.HandleFunc("PUT /v1/books/{book_id}", handlers.UpdateBook)
	mux.HandleFunc("DELETE /v1/books/{book_id}", handlers.DeleteBook)
}
