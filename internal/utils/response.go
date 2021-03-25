package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseSuccess returns passed payload in JSON format
func ResponseSuccess(w http.ResponseWriter, code int, payload interface{}) {
	jsonObj := map[string]interface{}{
		"success": true,
		"payload": payload,
	}

	response, _ := json.Marshal(jsonObj)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(response)
}

// ResponseError returns passed error message in JSON format
func ResponseError(w http.ResponseWriter, code int, message string) {
	jsonObj := map[string]interface{}{
		"success": false,
		"message": message,
	}

	response, _ := json.Marshal(jsonObj)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(response)
}
