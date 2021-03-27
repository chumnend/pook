package user

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Handler struct declaration
type Handler struct {
	DB *gorm.DB
}

// AttachHandler takes a router and adds routes to it
func AttachHandler(r *mux.Router, db *gorm.DB) {
	h := &Handler{DB: db}

	r.HandleFunc("/api/v1/register", h.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", h.Login).Methods("POST")

	r.HandleFunc("/api/v1/users", h.ListUsers).Methods("GET")
	r.HandleFunc("/api/v1/user/{id:[0-9]+}", h.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/user/{id:[0-9]+}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/v1/user/{id:[0-9]+}", h.DeleteUser).Methods("DELETE")
}

// Register creates a new user and returns jwt token
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - register")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// Login validates credentials and returns jwt token if valid
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - login")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// ListUsers returns list of Users found in the DB
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - list users")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// GetUser returns a User in the DB
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get user")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// UpdateUser updates the user in the DB
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update user")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// DeleteUser removes a user from the DB
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete user")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}
