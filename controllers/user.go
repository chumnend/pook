package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

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
		http.Error(w, "Empty request body", 400)
		return
	}

	// read content of request body into user struct
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Request must be JSON object", 400)
		return
	}

	// generate password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Something went wrong. Please try again later", 400)
		return
	}

	// save user information to db
	user.Password = string(password)
	createdUser := db.Create(&user)
	if createdUser.Error != nil {
		http.Error(w, "Something went wrong. Please try again later", 400)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

// Login - returns token if valid user
func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("endpoint: login \n")

	// validate request body
	if r.Body == nil {
		http.Error(w, "Empty request body", 400)
		return
	}

	// read content of request body into user struct
	credentials := models.User{}

	errJSON := json.NewDecoder(r.Body).Decode(&credentials)
	if errJSON != nil {
		http.Error(w, "Request must be JSON object", 400)
		return
	}

	// check if email exists
	foundUser := models.User{}
	errDB := db.Where("email = ?", credentials.Email).First(&foundUser).Error
	if errDB != nil {
		http.Error(w, "Invalid email and/or password", 400)
		return
	}

	// check if password is valid
	errPW := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credentials.Password))
	if errPW != nil && errPW == bcrypt.ErrMismatchedHashAndPassword {
		http.Error(w, "Invalid email and/or password", 400)
		return
	}

	// create token
	tk := &models.Token{
		ID:             foundUser.ID,
		Email:          foundUser.Email,
		StandardClaims: &jwt.StandardClaims{},
	}

	secret := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, errJWT := token.SignedString([]byte(secret))
	if errJWT != nil {
		http.Error(w, "Something went wrong. Please try again later", 400)
		return
	}

	json.NewEncoder(w).Encode(tokenString)
}
