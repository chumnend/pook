package main

import (
	"log"
	"os"

	"github.com/chumnend/pook/internal/pook"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("missing env: SECRET_KEY")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	s := pook.NewServer()
	s.Serve(":" + port)
}
