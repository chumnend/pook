package utils

import (
	"encoding/json"
	"net/http"
)

// SendJSON sends a JSON response with a given status code
func SendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}

// ParseJSON decodes a JSON request body into the given struct
func ParseJSON(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Prevent unknown fields from being parsed
	return decoder.Decode(v)
}
