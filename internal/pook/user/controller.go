package user

import (
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
	response.Error(w, http.StatusNotImplemented, "not yet implemented")
}

func (ctl *userCtl) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - login")
	response.Error(w, http.StatusNotImplemented, "not yet implemented")
}
