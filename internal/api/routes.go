package api

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", Pong)
}
