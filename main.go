package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type booking struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

var bookings []booking

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Bookings API: Ready to serve requests</p>")
	fmt.Println("Hit: home")
}

func getAllBookings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit: getAllBookings")
	json.NewEncoder(w).Encode(bookings)
}

func handleRequests() {
	// index routes
	http.HandleFunc("/", home)

	// booking routes
	http.HandleFunc("/v1/bookings", getAllBookings)

	log.Print("listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Preparing server...")

	bookings = []booking{
		booking{Title: "Test 1", Description: "This is a test booking"},
		booking{Title: "Test 2", Description: "This is a test booking"},
	}

	handleRequests()
}
