package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseSuccess returns passed payload in JSON format
func ResponseSuccess(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(response)
}

// ResponseError returns passed error message in JSON format
func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseSuccess(w, code, map[string]string{"error": message})
}
