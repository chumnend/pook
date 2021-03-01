package main

import (
	"log"
	"os"

	bookings "github.com/chumnend/bookings-server/server"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file was found")
	}

	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		log.Fatal("Missing env:  DATABASE_URL")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("Missing env:  SECRET")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Missing env:  PORT")
	}

	// run the server
	server := bookings.NewServer(connectionString, secret, port)
	server.Run()
}
