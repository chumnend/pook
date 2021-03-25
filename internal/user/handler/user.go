package handler

import (
	"fmt"
	"net/http"

	"github.com/chumnend/pook/internal/utils"

	"github.com/chumnend/pook/internal/user/domain"
	"github.com/gorilla/mux"
)

// UserHandler struct declaration
type UserHandler struct {
	User domain.UserEntity
}

// NewUserHandler creates new UserHandler
func NewUserHandler(r *mux.Router, user domain.UserEntity) *UserHandler {
	handler := &UserHandler{
		User: user,
	}

	r.HandleFunc("/api/v1/users/ping", handler.Ping).Methods("GET")
	r.HandleFunc("/api/v1/users", handler.FetchUsers).Methods("GET")
	r.HandleFunc("/api/v1/user/{id:[0-9]+}", handler.GetUserByID).Methods("GET")

	return handler
}

// Ping is a health check for the user routes
func (u *UserHandler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User API ready to serve requests")
}

// FetchUsers returns all users in db
func (u *UserHandler) FetchUsers(w http.ResponseWriter, r *http.Request) {
	utils.ResponseError(w, http.StatusNotImplemented, "Not yet implemented")
}

// GetUserByID returns a user by id passed in request
func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	utils.ResponseError(w, http.StatusNotImplemented, "Not yet implemented")
}
