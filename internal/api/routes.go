package api

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", pong)

	mux.HandleFunc("POST /v1/register", register)
	mux.HandleFunc("POST /v1/login", login)
	mux.HandleFunc("GET /v1/users/{user_id}", getUser)
}
