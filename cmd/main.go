package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/chumnend/pook/internal/books"
	"github.com/chumnend/pook/internal/comments"
	"github.com/chumnend/pook/internal/pages"
	"github.com/chumnend/pook/internal/ratings"
	"github.com/chumnend/pook/internal/users"
	"github.com/chumnend/pook/internal/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("PG_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}
	log.Println("Connected to the database successfully")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", utils.Pong)
	users.RegisterRoutes(mux)
	books.RegisterRoutes(mux, db)
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
