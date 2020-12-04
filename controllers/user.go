package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/chumnend/bookings-server/models"
	"github.com/chumnend/bookings-server/utils"
)

var db = utils.ConnectDB()

// Register - creates a new user
func Register(w http.ResponseWriter, r *http.Request) {
	log.Printf("endpoint: register \n")

	// validate request body
	if r.Body == nil {
		http.Error(w, "Invalid: Empty request body", 400)
		return
	}

	// read content of request body into user struct
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid: Request must be JSON object", 400)
		return
	}

	// generate password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Invalid: Something went wrong. Please try again later", 400)
		return
	}

	// save user information to db
	user.Password = string(password)
	createdUser := db.Create(&user)
	if createdUser.Error != nil {
		http.Error(w, "Invalid: Something went wrong. Please try again later", 400)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

// Login - returns token if valid user
func Login(w http.ResponseWriter, r *http.Request) {
	// read content of request body into user struct
	// check if email and password are valid
	// create token
}
