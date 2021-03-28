package user

import (
	"encoding/json"
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

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")

	r.HandleFunc("/users", h.ListUsers).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", h.GetUser).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id:[0-9]+}", h.DeleteUser).Methods("DELETE")
}

// Register creates a new user and returns jwt token
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// create new user struct
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// call method to create user in DB
	if err := u.Create(h.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// generate jwt token
	if token, err := u.GenerateToken(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
	}
}

// Login validates credentials and returns jwt token if valid
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// get credentials from request
	var creds User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// check to see if user exists
	u, err := FindUserByEmail(h.DB, creds.Email)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Email and/or Password")
		return
	}

	// compare password with found users
	isValid := u.comparePassword(creds.Password)
	if !isValid {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Email and/or Password")
	}

	// generate jwt token
	if token, err := u.GenerateToken(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
	}
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
