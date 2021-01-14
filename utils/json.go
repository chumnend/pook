package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse struct declaration
type JSONResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// SendJSONResponse - sends json response
func SendJSONResponse(w http.ResponseWriter, response JSONResponse, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}
