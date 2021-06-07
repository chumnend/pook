package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/response"
)

// Controller interface declaration
type Controller interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type userCtl struct {
	srv Service
}

// NewUserController creates a Controller with given service
func NewUserController(srv Service) Controller {
	return &userCtl{srv: srv}
}

func (ctl *userCtl) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - register")

	// create new user struct
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// validate the user struct
	validateErr := ctl.srv.Validate(&u)
	if validateErr != nil {
		response.Error(w, http.StatusBadRequest, "missing and/or invalid information")
		return
	}

	// call method to create user in DB
	if err := ctl.srv.Save(&u); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// generate jwt token
	if token, err := ctl.srv.GenerateToken(&u); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{"token": token})
	}
}

func (ctl *userCtl) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - login")

	// get credentials from request
	var creds User
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
