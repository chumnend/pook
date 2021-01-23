package bookings

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

// Send the JSON response
func (response *JSONResponse) Send(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}
