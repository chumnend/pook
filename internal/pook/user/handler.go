package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/response"
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

	r.HandleFunc("/register", h.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", h.Login).Methods("POST", "OPTIONS")

	r.HandleFunc("/users", h.ListUsers).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", h.GetUser).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id:[0-9]+}", h.DeleteUser).Methods("DELETE")
}

// Register creates a new user and returns jwt token
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - register")

	// create new user struct
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// validate the user struct
	isValid := u.Validate()
	if !isValid {
		response.Error(w, http.StatusBadRequest, "missing and/or invalid information")
		return
	}

	// call method to create user in DB
	if err := u.Create(h.DB); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// generate jwt token
	if token, err := u.GenerateToken(); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{"token": token})
	}
}

// Login validates credentials and returns jwt token if valid
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - login")

	// get credentials from request
	var creds User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// check to see if user exists
	u, err := FindUserByEmail(h.DB, creds.Email)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid email and/or password")
		return
	}

	// compare password with found users
	isValid := u.ComparePassword(creds.Password)
	if !isValid {
		response.Error(w, http.StatusBadRequest, "invalid email and/or password")
		return
	}

	// generate jwt token
	if token, err := u.GenerateToken(); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{"token": token})
	}
}

// ListUsers returns list of Users found in the DB
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - list users")
	response.Error(w, http.StatusNotImplemented, "not yet implemented")
}

// GetUser returns a User in the DB
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get user")
	response.Error(w, http.StatusNotImplemented, "not yet implemented")
}

// UpdateUser updates the user in the DB
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update user")
	response.Error(w, http.StatusNotImplemented, "not yet implemented")
}

// DeleteUser removes a user from the DB
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete user")
	response.Error(w, http.StatusNotImplemented, "not yet implemented")
}
