package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/chumnend/bookings-server/internal/bookings"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("DATABASE_URL")
	secret := os.Getenv("SECRET_KEY")
	port := os.Getenv("PORT")

	app := bookings.NewApp()
	app.Init(connectionString, secret)
	app.Run(":" + port)
}
