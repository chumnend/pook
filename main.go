package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type app struct {
	router *mux.Router
	db     *gorm.DB
}

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	a := new(app)

	a.db, err = gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	a.router = mux.NewRouter()
	a.router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ready to serve requests")
	})

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, a.router))
}
