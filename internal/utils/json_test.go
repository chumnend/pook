package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendJSON(t *testing.T) {
	// Mock response writer
	rr := httptest.NewRecorder()

	// Test data
	data := map[string]string{"message": "success"}
	statusCode := http.StatusOK

	// Call SendJSON
	SendJSON(rr, data, statusCode)

	// Check status code
	if rr.Code != statusCode {
		t.Errorf("expected status code %d, got %d", statusCode, rr.Code)
	}

	// Check content type
	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", rr.Header().Get("Content-Type"))
	}

	// Check response body
	var responseData map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &responseData)
	if err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if responseData["message"] != "success" {
		t.Errorf("expected message 'success', got '%s'", responseData["message"])
	}
}

func TestParseJSON(t *testing.T) {
	// Test data
	data := map[string]string{"key": "value"}
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("failed to marshal test data: %v", err)
	}

	// Mock request
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Call ParseJSON
	var parsedData map[string]string
	err = ParseJSON(req, &parsedData)
	if err != nil {
		t.Fatalf("ParseJSON failed: %v", err)
	}

	// Check parsed data
	if parsedData["key"] != "value" {
		t.Errorf("expected key 'value', got '%s'", parsedData["key"])
	}
}

func TestParseJSONWithInvalidBody(t *testing.T) {
	// Invalid JSON body
	body := []byte(`{"key": "value",}`)

	// Mock request
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Call ParseJSON
	var parsedData map[string]string
	err := ParseJSON(req, &parsedData)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}
