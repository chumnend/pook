package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chumnend/bookings-server/routes"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Preparing server...")

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup routes
	routes.HandleRequests()

	// listen
	port := os.Getenv("PORT")
	log.Printf("listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
