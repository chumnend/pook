package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Booking ...
type Booking struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Bookings ...
var Bookings []Booking

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Bookings API: Ready to serve requests</p>")

	fmt.Println("Hit: home")
}

func getAllBookings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Bookings)

	fmt.Println("Hit: getAllBookings")
}

func createBooking(w http.ResponseWriter, r *http.Request) {
	var booking Booking
	json.NewDecoder(r.Body).Decode(&booking)

	Bookings = append(Bookings, booking)

	json.NewEncoder(w).Encode(booking)
}

func getBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, b := range Bookings {
		if b.ID == id {
			json.NewEncoder(w).Encode(b)
		}
	}

	fmt.Printf("Hit: getBooking %s\n", id)
}

func updateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, b := range Bookings {
		if b.ID == id {
			var booking Booking
			json.NewDecoder(r.Body).Decode(&booking)

			Bookings[i] = booking
		}
	}

	fmt.Printf("Hit: updateBooking %s\n", id)
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, b := range Bookings {
		if b.ID == id {
			Bookings = append(Bookings[:i], Bookings[i+1:]...)
		}
	}

	fmt.Printf("Hit: deleteBooking %s\n", id)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	// index routes
	router.HandleFunc("/", home)

	// booking routes
	router.HandleFunc("/v1/bookings", getAllBookings).Methods("GET")
	router.HandleFunc("/v1/bookings", createBooking).Methods("POST")
	router.HandleFunc("/v1/bookings/{id}", getBooking).Methods("GET")
	router.HandleFunc("/v1/bookings/{id}", updateBooking).Methods("PUT")
	router.HandleFunc("/v1/bookings/{id}", deleteBooking).Methods("DELETE")

	// listen
	log.Print("listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	fmt.Println("Preparing server...")

	// setup sample data
	Bookings = []Booking{
		Booking{ID: "1", Title: "Test 1", Description: "This is a test booking"},
		Booking{ID: "2", Title: "Test 2", Description: "This is a test booking"},
	}

	handleRequests()
}
