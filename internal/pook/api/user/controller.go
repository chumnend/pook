package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/pook/domain"
	"github.com/chumnend/pook/internal/pook/response"
)

type userCtl struct {
	srv domain.UserService
}

// NewController creates a UserController with given UserService
func NewController(srv domain.UserService) domain.UserController {
	return &userCtl{srv: srv}
}

func (ctl *userCtl) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - register")

	// create new user struct
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// validate the user struct
	validateErr := ctl.srv.Validate(&user)
	if validateErr != nil {
		response.Error(w, http.StatusBadRequest, "missing and/or invalid information")
		return
	}

	// call method to create user in DB
	if err := ctl.srv.Save(&user); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// generate jwt token
	if token, err := ctl.srv.GenerateToken(&user); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{"token": token})
	}
}

func (ctl *userCtl) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - login")

	// get credentials from request
	var creds domain.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// check to see if user exists
	u, err := ctl.srv.FindByEmail(creds.Email)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid email and/or password")
		return
	}

	// compare password with found users
	pwErr := ctl.srv.ComparePassword(u, creds.Password)
	if pwErr != nil {
		response.Error(w, http.StatusBadRequest, "invalid email and/or password")
		return
	}

	// generate jwt token
	if token, err := ctl.srv.GenerateToken(u); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{"token": token})
	}
}
