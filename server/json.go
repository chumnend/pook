package server

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

// SendError writes an error json message to client
func SendError(w http.ResponseWriter, code int, message string) {
	SendJSON(w, code, JSONResponse{
		Success: false,
		Message: message,
	})
}

// SendJSON writes a JSON payload back to client
func SendJSON(w http.ResponseWriter, code int, payload JSONResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(payload)
}
