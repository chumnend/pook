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
	router.HandleFunc("/", controllers.Ready)

	// listen
	log.Print("listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
