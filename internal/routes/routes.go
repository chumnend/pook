package routes

import (
	"net/http"

	"github.com/chumnend/pook/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Health Check Route
	mux.HandleFunc("GET /v1/status", handlers.Ping)
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
	// Page Routes
	mux.HandleFunc("POST /books/{book_id}/pages", handlers.CreatePage)
	mux.HandleFunc("GET /books/{book_id}/pages", handlers.GetPages)
	mux.HandleFunc("GET /books/{book_id}/pages/{page_id}", handlers.GetPage)
	mux.HandleFunc("PUT /books/{book_id}/pages/{page_id}", handlers.UpdatePage)
	mux.HandleFunc("DELETE /books/{book_id}/pages/{page_id}", handlers.DeletePage)
}
