package main

import (
	"log"
	"os"

	"github.com/chumnend/pook/web"
	"github.com/joho/godotenv"
)

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
	app := web.NewApp(dbURL, port)
	app.Start()
}
