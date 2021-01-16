package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	bookings "github.com/chumnend/bookings-server/internal"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	app := bookings.App{}
	app.Init(connectionString)
	app.Run(":" + port)
}
