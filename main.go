package main

import (
	"log"
	"os"

	"github.com/chumnend/pook/server"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	// load environment variables
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
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
	s := server.New()
	s.Initialize(dbURL, port)
	s.Run()
}
