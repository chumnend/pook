package response

import (
	"encoding/json"
	"net/http"
)

// JSON passes a JSON message to ResponseWriter struct
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(payload)
}

// Error passes a JSON message with error to ResponseWriter struct
func Error(w http.ResponseWriter, code int, message string) {
	JSON(w, code, map[string]string{"error": message})
}
