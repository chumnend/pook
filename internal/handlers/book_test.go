package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chumnend/pook/internal/db"
	"github.com/google/uuid"
)

func TestCreateBook(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	originalDB := db.DB
	db.DB = mockDB
	defer func() {
		db.DB = originalDB
	}()

	userID := uuid.New()
	requestBody := map[string]interface{}{
		"userId":   userID.String(),
		"imageUrl": "https://example.com/book-cover.jpg",
		"title":    "Test Book Title",
	}

	mock.ExpectExec("INSERT INTO books").
		WithArgs(sqlmock.AnyArg(), userID, "https://example.com/book-cover.jpg", "Test Book Title", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CreateBook returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("CreateBook returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	expectedMessage := "book successfully created"
	if response["message"] != expectedMessage {
		t.Errorf("CreateBook returned wrong message: got %v want %v", response["message"], expectedMessage)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestCreateBookInvalidInput(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/books", strings.NewReader("invalid json"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("CreateBook returned wrong status code for invalid JSON: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid input") {
		t.Errorf("CreateBook should return 'invalid input' error message")
	}
}

func TestCreateBookMissingRequiredFields(t *testing.T) {
	testCases := []struct {
		name string
		body map[string]interface{}
	}{
		{
			name: "missing title",
			body: map[string]interface{}{
				"userId":   uuid.New().String(),
				"imageUrl": "https://example.com/image.jpg",
			},
		},
		{
			name: "missing userId",
			body: map[string]interface{}{
				"imageUrl": "https://example.com/image.jpg",
				"title":    "Test Title",
			},
		},
		{
			name: "missing imageUrl",
			body: map[string]interface{}{
				"userId": uuid.New().String(),
				"title":  "Test Title",
			},
		},
		{
			name: "empty title",
			body: map[string]interface{}{
				"userId":   uuid.New().String(),
				"imageUrl": "https://example.com/image.jpg",
				"title":    "",
			},
		},
		{
			name: "empty userId",
			body: map[string]interface{}{
				"userId":   "",
				"imageUrl": "https://example.com/image.jpg",
				"title":    "Test Title",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tc.body)
			req, err := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CreateBook)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("CreateBook returned wrong status code for %s: got %v want %v", tc.name, status, http.StatusBadRequest)
			}

			responseBody := rr.Body.String()
			if !strings.Contains(responseBody, "required") {
				t.Errorf("CreateBook should return error message about required fields for %s, got: %s", tc.name, responseBody)
			}
		})
	}
}

func TestCreateBookInvalidUserId(t *testing.T) {
	requestBody := map[string]interface{}{
		"userId":   "invalid-uuid",
		"imageUrl": "https://example.com/book-cover.jpg",
		"title":    "Test Book Title",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("CreateBook returned wrong status code for invalid UUID: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid user id") {
		t.Errorf("CreateBook should return 'invalid user id' error message")
	}
}

func TestCreateBookCreationErr(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	originalDB := db.DB
	db.DB = mockDB
	defer func() {
		db.DB = originalDB
	}()

	userID := uuid.New()
	requestBody := map[string]interface{}{
		"userId":   userID.String(),
		"imageUrl": "https://example.com/book-cover.jpg",
		"title":    "Test Book Title",
	}

	mock.ExpectExec("INSERT INTO books").
		WithArgs(sqlmock.AnyArg(), userID, "https://example.com/book-cover.jpg", "Test Book Title", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(sqlmock.ErrCancelled)

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("CreateBook returned wrong status code for DB error: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "unable to create book") {
		t.Errorf("CreateBook should return 'unable to create book' error message, got: %s", rr.Body.String())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}
