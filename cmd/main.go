package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chumnend/pook/internal/books"
	"github.com/chumnend/pook/internal/comments"
	"github.com/chumnend/pook/internal/pages"
	"github.com/chumnend/pook/internal/ratings"
	"github.com/chumnend/pook/internal/users"
	"github.com/chumnend/pook/internal/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", utils.Pong)
	users.RegisterRoutes(mux)
	books.RegisterRoutes(mux)
	pages.RegisterRoutes(mux)
	comments.RegisterRoutes(mux)
	ratings.RegisterRoutes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on port", port)
	http.ListenAndServe(":"+port, mux)
}
