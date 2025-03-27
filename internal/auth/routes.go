package auth

import (
	"net/http"

	"github.com/chumnend/pook/internal/utils"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/register", utils.NotImplemented)
	mux.HandleFunc("POST /v1/login", utils.NotImplemented)
}
