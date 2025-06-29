package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chumnend/pook/internal/db"
	"github.com/google/uuid"
)

func TestCreatePage(t *testing.T) {
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

	bookID := uuid.New()
	requestBody := map[string]interface{}{
		"imageUrl":  "https://example.com/page-image.jpg",
		"caption":   "Test page caption",
		"pageOrder": 1,
	}

	mock.ExpectExec("INSERT INTO pages").
		WithArgs(sqlmock.AnyArg(), bookID, "https://example.com/page-image.jpg", "Test page caption", 1, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", fmt.Sprintf("/v1/books/%s/pages", bookID.String()), bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("book_id", bookID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CreatePage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("CreatePage returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	expectedMessage := "page successfully created"
	if response["message"] != expectedMessage {
		t.Errorf("CreatePage returned wrong message: got %v want %v", response["message"], expectedMessage)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestCreatePageInvalidInput(t *testing.T) {
	bookID := uuid.New()
	req, err := http.NewRequest("POST", fmt.Sprintf("/v1/books/%s/pages", bookID.String()), strings.NewReader("invalid json"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("book_id", bookID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("CreatePage returned wrong status code for invalid JSON: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid input") {
		t.Errorf("CreatePage should return 'invalid input' error message")
	}
}

func TestCreatePageMissingRequiredFields(t *testing.T) {
	testCases := []struct {
		name string
		body map[string]interface{}
	}{
		{
			name: "missing imageUrl",
			body: map[string]interface{}{
				"caption":   "Test caption",
				"pageOrder": 1,
			},
		},
		{
			name: "missing caption",
			body: map[string]interface{}{
				"imageUrl":  "https://example.com/image.jpg",
				"pageOrder": 1,
			},
		},
		{
			name: "negative pageOrder",
			body: map[string]interface{}{
				"imageUrl":  "https://example.com/image.jpg",
				"caption":   "Test caption",
				"pageOrder": -1,
			},
		},
		{
			name: "empty imageUrl",
			body: map[string]interface{}{
				"imageUrl":  "",
				"caption":   "Test caption",
				"pageOrder": 1,
			},
		},
		{
			name: "empty caption",
			body: map[string]interface{}{
				"imageUrl":  "https://example.com/image.jpg",
				"caption":   "",
				"pageOrder": 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bookID := uuid.New()
			jsonBody, _ := json.Marshal(tc.body)
			req, err := http.NewRequest("POST", fmt.Sprintf("/v1/books/%s/pages", bookID.String()), bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetPathValue("book_id", bookID.String())

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CreatePage)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("CreatePage should return 400 for %s: got %v want %v", tc.name, status, http.StatusBadRequest)
			}

			expectedError := "all fields (bookId, imageUrl, caption, pageOrder) are required"
			if !strings.Contains(rr.Body.String(), expectedError) {
				t.Errorf("CreatePage should return validation error for %s", tc.name)
			}
		})
	}
}

func TestCreatePageInvalidBookID(t *testing.T) {
	requestBody := map[string]interface{}{
		"imageUrl":  "https://example.com/page-image.jpg",
		"caption":   "Test page caption",
		"pageOrder": 1,
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/books/invalid-uuid/pages", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("book_id", "invalid-uuid")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("CreatePage returned wrong status code for invalid book ID: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid book id") {
		t.Errorf("CreatePage should return 'invalid book id' error message")
	}
}

func TestGetPages(t *testing.T) {
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

	bookID := uuid.New()
	pageID1 := uuid.New()
	pageID2 := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "book_id", "image_url", "caption", "page_order", "created_at", "updated_at"}).
		AddRow(pageID1, bookID, "https://example.com/page1.jpg", "Page 1 caption", 1, now, now).
		AddRow(pageID2, bookID, "https://example.com/page2.jpg", "Page 2 caption", 2, now, now)

	mock.ExpectQuery("SELECT id, book_id, image_url, caption, page_order, created_at, updated_at FROM pages WHERE book_id = \\$1 ORDER BY page_order").
		WithArgs(bookID).
		WillReturnRows(rows)

	req, err := http.NewRequest("GET", fmt.Sprintf("/v1/books/%s/pages", bookID.String()), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("book_id", bookID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPages)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetPages returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("GetPages returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	pages, ok := response["pages"].([]interface{})
	if !ok {
		t.Errorf("Response should contain pages array")
	}

	if len(pages) != 2 {
		t.Errorf("Expected 2 pages, got %d", len(pages))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestGetPagesInvalidBookID(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/books/invalid-uuid/pages", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("book_id", "invalid-uuid")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPages)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("GetPages returned wrong status code for invalid book ID: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid book id") {
		t.Errorf("GetPages should return 'invalid book id' error message")
	}
}

func TestGetPage(t *testing.T) {
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

	pageID := uuid.New()
	bookID := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "book_id", "image_url", "caption", "page_order", "created_at", "updated_at"}).
		AddRow(pageID, bookID, "https://example.com/page.jpg", "Test caption", 1, now, now)

	mock.ExpectQuery("SELECT id, book_id, image_url, caption, page_order, created_at, updated_at FROM pages WHERE id = \\$1").
		WithArgs(pageID).
		WillReturnRows(rows)

	req, err := http.NewRequest("GET", fmt.Sprintf("/v1/pages/%s", pageID.String()), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("page_id", pageID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetPage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("GetPage returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	page, ok := response["page"].(map[string]interface{})
	if !ok {
		t.Errorf("Response should contain page object")
	}

	if page["caption"] != "Test caption" {
		t.Errorf("Expected caption 'Test caption', got %v", page["caption"])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestGetPageInvalidPageID(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/pages/invalid-uuid", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("page_id", "invalid-uuid")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("GetPage returned wrong status code for invalid page ID: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid page id") {
		t.Errorf("GetPage should return 'invalid page id' error message")
	}
}

func TestUpdatePage(t *testing.T) {
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

	pageID := uuid.New()
	requestBody := map[string]interface{}{
		"imageUrl":  "https://example.com/updated-page.jpg",
		"caption":   "Updated page caption",
		"pageOrder": 2,
	}

	mock.ExpectExec("UPDATE pages SET image_url = \\$1, caption = \\$2, page_order = \\$3, updated_at = \\$4 WHERE id = \\$5").
		WithArgs("https://example.com/updated-page.jpg", "Updated page caption", 2, sqlmock.AnyArg(), pageID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("PUT", fmt.Sprintf("/v1/pages/%s", pageID.String()), bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("page_id", pageID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdatePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("UpdatePage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("UpdatePage returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	expectedMessage := "page successfully updated"
	if response["message"] != expectedMessage {
		t.Errorf("UpdatePage returned wrong message: got %v want %v", response["message"], expectedMessage)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestUpdatePageInvalidInput(t *testing.T) {
	pageID := uuid.New()
	req, err := http.NewRequest("PUT", fmt.Sprintf("/v1/pages/%s", pageID.String()), strings.NewReader("invalid json"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("page_id", pageID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdatePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("UpdatePage returned wrong status code for invalid JSON: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid input") {
		t.Errorf("UpdatePage should return 'invalid input' error message")
	}
}

func TestUpdatePageMissingRequiredFields(t *testing.T) {
	testCases := []struct {
		name string
		body map[string]interface{}
	}{
		{
			name: "missing imageUrl",
			body: map[string]interface{}{
				"caption":   "Test caption",
				"pageOrder": 1,
			},
		},
		{
			name: "missing caption",
			body: map[string]interface{}{
				"imageUrl":  "https://example.com/image.jpg",
				"pageOrder": 1,
			},
		},
		{
			name: "negative pageOrder",
			body: map[string]interface{}{
				"imageUrl":  "https://example.com/image.jpg",
				"caption":   "Test caption",
				"pageOrder": -1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pageID := uuid.New()
			jsonBody, _ := json.Marshal(tc.body)
			req, err := http.NewRequest("PUT", fmt.Sprintf("/v1/pages/%s", pageID.String()), bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetPathValue("page_id", pageID.String())

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(UpdatePage)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("UpdatePage should return 400 for %s: got %v want %v", tc.name, status, http.StatusBadRequest)
			}

			expectedError := "all fields (imageUrl, caption, pageOrder) are required"
			if !strings.Contains(rr.Body.String(), expectedError) {
				t.Errorf("UpdatePage should return validation error for %s", tc.name)
			}
		})
	}
}

func TestUpdatePageInvalidPageID(t *testing.T) {
	requestBody := map[string]interface{}{
		"imageUrl":  "https://example.com/updated-page.jpg",
		"caption":   "Updated page caption",
		"pageOrder": 2,
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("PUT", "/v1/pages/invalid-uuid", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("page_id", "invalid-uuid")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdatePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("UpdatePage returned wrong status code for invalid page ID: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid page id") {
		t.Errorf("UpdatePage should return 'invalid page id' error message")
	}
}

func TestDeletePage(t *testing.T) {
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

	pageID := uuid.New()

	mock.ExpectExec("DELETE FROM pages WHERE id = \\$1").
		WithArgs(pageID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/v1/pages/%s", pageID.String()), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("page_id", pageID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeletePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("DeletePage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("DeletePage returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	expectedMessage := "page successfully deleted"
	if response["message"] != expectedMessage {
		t.Errorf("DeletePage returned wrong message: got %v want %v", response["message"], expectedMessage)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestDeletePageInvalidPageID(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/pages/invalid-uuid", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("page_id", "invalid-uuid")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeletePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("DeletePage returned wrong status code for invalid page ID: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid page id") {
		t.Errorf("DeletePage should return 'invalid page id' error message")
	}
}
