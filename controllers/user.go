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
		resp := utils.JSONResponse{
			Success: false,
			Message: "Empty request body",
		}

		utils.SendJSONResponse(w, resp, 400)
		return
	}

	// read content of request body into user struct
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp := utils.JSONResponse{
			Success: false,
			Message: "Request must be JSON object",
		}

		utils.SendJSONResponse(w, resp, 400)
		return
	}

	// generate password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		resp := utils.JSONResponse{
			Success: false,
			Message: "Something went wrong. Please try again later",
		}

		utils.SendJSONResponse(w, resp, 400)
		return
	}

	// save user information to db
	user.Password = string(password)
	result := db.Create(&user)
	if result.Error != nil {
		resp := utils.JSONResponse{
			Success: false,
			Message: "Email already exists",
		}

		utils.SendJSONResponse(w, resp, 400)
		return
	}

	// create token
	tk := &models.Token{
		ID:             user.ID,
		Email:          user.Email,
		StandardClaims: &jwt.StandardClaims{},
	}

	secret := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, errJWT := token.SignedString([]byte(secret))
	if errJWT != nil {
		resp := utils.JSONResponse{
			Success: false,
			Message: "Something went wrong. Please try again later",
		}

		utils.SendJSONResponse(w, resp, 400)
		return
	}

	resp := utils.JSONResponse{
		Success: true,
		Message: "New user created",
		Payload: map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
			"token": tokenString,
		},
	}

	utils.SendJSONResponse(w, resp, 200)
}

// Login - returns token if valid user
func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("endpoint: login \n")

	// validate request body
	if r.Body == nil {
		resp := utils.JSONResponse{
			Success: false,
			Message: "Empty request body",
		}

		utils.SendJSONResponse(w, resp, 400)
		return
	}

	// read content of request body into user struct
	credentials := models.User{}

	errJSON := json.NewDecoder(r.Body).Decode(&credentials)
	if errJSON != nil {
		resp := utils.JSONResponse{
			Success: false,
			Message: "Request must be JSON object",
		}

		utils.SendJSONResponse(w, resp, 400)
		return
	}

	// check if email exists
	foundUser := models.User{}
	errDB := db.Where("email = ?", credentials.Email).First(&foundUser).Error
	if errDB != nil {
		resp := utils.JSONResponse{
			Success: false,
			Message: "Invalid email and/or password",
		}

		utils.SendJSONResponse(w, resp, 400)
		return
	}

	// check if password is valid
	errPW := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credentials.Password))
	if errPW != nil && errPW == bcrypt.ErrMismatchedHashAndPassword {
		resp := utils.JSONResponse{
			Success: false,
			Message: "Invalid email and/or password",
		}

		utils.SendJSONResponse(w, resp, 400)
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
		resp := utils.JSONResponse{
			Success: false,
			Message: "Something went wrong. Please try again later",
		}

		utils.SendJSONResponse(w, resp, 400)
		return
	}

	resp := utils.JSONResponse{
		Success: true,
		Message: "New user created",
		Payload: map[string]interface{}{
			"id":    foundUser.ID,
			"email": foundUser.Email,
			"token": tokenString,
		},
	}

	utils.SendJSONResponse(w, resp, 200)
}
