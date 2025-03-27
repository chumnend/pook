package ratings

import (
	"net/http"

	"github.com/chumnend/pook/internal/utils"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/books/{book_id}/ratings/", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/ratings/", utils.NotImplemented)
	mux.HandleFunc("GET /v1/books/{book_id}/ratings/{rating_id}", utils.NotImplemented)
	mux.HandleFunc("PUT /v1/books/{book_id}/ratings/{rating_id}", utils.NotImplemented)
	mux.HandleFunc("DELETE /v1/books/{book_id}/ratings/{rating_id}", utils.NotImplemented)
}
