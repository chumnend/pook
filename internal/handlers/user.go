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
		response := map[string]string{
			"message": "invalid input",
		}
		utils.SendJSON(w, response, http.StatusBadRequest)
		return
	}

	if input.Email == "" || input.Username == "" || input.Password == "" {
		response := map[string]string{
			"message": "all fields (email, username, password) are required",
		}
		utils.SendJSON(w, response, http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(input.Email) {
		response := map[string]string{
			"message": "invalid email format",
		}
		utils.SendJSON(w, response, http.StatusBadRequest)
		return
	}

	if err := models.CreateUser(input.Username, input.Email, input.Password); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		utils.SendJSON(w, response, http.StatusInternalServerError)
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
		response := map[string]string{
			"message": "invalid input",
		}
		utils.SendJSON(w, response, http.StatusBadRequest)
		return
	}

	if input.Username == "" || input.Password == "" {
		response := map[string]string{
			"message": "all fields (username, password) are required",
		}
		utils.SendJSON(w, response, http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByUsername(input.Username)
	if err != nil {
		response := map[string]string{
			"message": "invalid username and/or password",
		}
		utils.SendJSON(w, response, http.StatusBadRequest)
		return
	}

	if err := models.ComparePassword(user, input.Password); err != nil {
		response := map[string]string{
			"message": "invalid username and/or password",
		}
		utils.SendJSON(w, response, http.StatusBadRequest)
		return
	}

	token, err := models.GenerateUserToken(user)
	if err != nil {
		response := map[string]string{
			"message": "something went wrong",
		}
		utils.SendJSON(w, response, http.StatusInternalServerError)
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
		response := map[string]string{
			"message": "invalid user id",
		}
		utils.SendJSON(w, response, http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByID(parsed_uuid)
	if err != nil {
		response := map[string]string{
			"message": "user not found",
		}
		utils.SendJSON(w, response, http.StatusBadRequest)
		return
	}

	utils.SendJSON(w, user, http.StatusFound)
}
