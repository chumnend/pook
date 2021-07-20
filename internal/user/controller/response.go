package controller

import (
	"encoding/json"
	"net/http"
)

// responWithJSON passes a JSON message to ResponseWriter struct
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(payload)
}

// SendError passes a JSON message with error to ResponseWriter struct
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
