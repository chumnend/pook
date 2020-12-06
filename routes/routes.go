package routes

import (
	"github.com/gorilla/mux"

	"github.com/chumnend/bookings-server/controllers"
)

// HandleRequests - handle requests
func HandleRequests() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// index routes
	router.HandleFunc("/", controllers.Ready).Methods("GET")

	// auth routes
	router.HandleFunc("/api/v1/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")

	return router
}
