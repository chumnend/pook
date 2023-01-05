package controller

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func checkErrorMessage(t *testing.T, rr *httptest.ResponseRecorder) {
	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	value, ok := m["error"]
	if !ok {
		t.Errorf("Unable to find key 'error'.")
	}
	if value == "" {
		t.Errorf("Expected 'error' to be non empty. Got %v.", value)
	}
}
