package handlers

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/models"
	"github.com/chumnend/pook/internal/utils"
	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, req *http.Request) {
	log.Println("Request made to" + req.URL.Path)

	type requestInput struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input requestInput
	if err := utils.ParseJSON(req, &input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if input.Email == "" || input.Username == "" || input.Password == "" {
		http.Error(w, "all fields (email, username, password) are required", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(input.Email) {
		http.Error(w, "invalid email format", http.StatusBadRequest)
		return
	}

	if err := models.CreateUser(input.Username, input.Email, input.Password); err != nil {
		http.Error(w, "unable to create user", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"message": "registration successful",
	}
	utils.SendJSON(w, response, http.StatusOK)
}

func Login(w http.ResponseWriter, req *http.Request) {
	log.Println("Request made to" + req.URL.Path)

	type requestInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input requestInput
	if err := utils.ParseJSON(req, &input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if input.Username == "" || input.Password == "" {
		http.Error(w, "all fields (username, password) are required", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByUsername(input.Username)
	if err != nil {
		http.Error(w, "invalid username and/or password", http.StatusBadRequest)
		return
	}

	if err := models.ComparePassword(user, input.Password); err != nil {
		http.Error(w, "invalid username and/or password", http.StatusBadRequest)
		return
	}

	token, err := models.GenerateUserToken(user)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"id":       user.ID.String(),
		"email":    user.Email,
		"username": user.Username,
		"token":    token,
	}

	utils.SendJSON(w, response, http.StatusOK)
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	log.Println("Request made to" + req.URL.Path)

	user_id := req.PathValue("user_id")

	parsed_uuid, err := uuid.Parse(user_id)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByID(parsed_uuid)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	utils.SendJSON(w, user, http.StatusFound)
}
