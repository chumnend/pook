package user

import (
	"context"
	"net/http"

	"github.com/chumnend/pook/internal/utils"
	"github.com/gorilla/mux"
)

// UserHandler struct declaration
type UserHandler struct {
	User UserUsecase
}

// AddUserHandler creates new UserHandler
func AddUserHandler(r *mux.Router, user UserUsecase) *UserHandler {
	handler := &UserHandler{
		User: user,
	}

	r.HandleFunc("/api/v1/users", handler.FetchUsers).Methods("GET")
	r.HandleFunc("/api/v1/user/{id:[0-9]+}", handler.GetUserByID).Methods("GET")

	return handler
}

// FetchUsers returns all users in db
func (u *UserHandler) FetchUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.User.Fetch(context.TODO())
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, users)
}

// GetUserByID returns a user by id passed in request
func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	user, err := u.User.GetByID(context.TODO(), id)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, "User not found")
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, user)
}
