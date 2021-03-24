package handler

import (
	"fmt"
	"net/http"

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

	r.HandleFunc("/api/v1/users/ping", handler.Ping)

	return handler
}

// Ping is a health check for the user routes
func (u *UserHandler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User API ready to serve requests")
}
