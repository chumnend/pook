package users

import (
	"net/http"

	"github.com/chumnend/pook/internal/utils"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/register", utils.NotImplemented)
	mux.HandleFunc("POST /v1/login", utils.NotImplemented)

	mux.HandleFunc("GET /v1/users/{user_id}", utils.NotImplemented)
}
