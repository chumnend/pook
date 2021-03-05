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
	addr   string
	router *mux.Router
	db     *gorm.DB
}

func (a *app) setup(dbURL string, port string) {
	var err error

	// connect database
	a.db, err = gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// setup api routes
	a.router = mux.NewRouter().StrictSlash(true)
	a.router.HandleFunc("/status", statusHandler)

	a.addr = ":" + port
}

func (a *app) start() {
	log.Println("Listening on " + a.addr)
	log.Fatal(http.ListenAndServe(a.addr, a.router))
}

func main() {
	var err error

	// load environment variables
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

	// create app instance
	a := new(app)
	a.setup(dbURL, port)
	a.start()
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready to serve requests")
}
