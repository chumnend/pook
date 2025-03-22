package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPong(t *testing.T) {
	req, err := http.NewRequest("GET", "/pong", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Pong)

	handler.ServeHTTP(rr, req)

	expected := "Pong\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
