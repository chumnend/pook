package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON passes a JSON message to ResponseWriter struct
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(payload)
}

// RespondWithError passes a JSON message with error to ResponseWriter struct
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
