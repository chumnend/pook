package books

import (
	"database/sql"
	"net/http"

	"github.com/chumnend/pook/internal/utils"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("POST /v1/books", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}", utils.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}", utils.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}", utils.NotImplemented)
}
