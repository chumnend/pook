package routes

import (
	"log"
	"net/http"

	"github.com/chumnend/bookings-server/controllers"
	"github.com/gorilla/mux"
)

// HandleRequests - handle requests
func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	// index routes
	router.HandleFunc("/", controllers.Ready).Methods("GET")

	// auth routes
	router.HandleFunc("/api/v1/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")

	// listen
	log.Print("listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
