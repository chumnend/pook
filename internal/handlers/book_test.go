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
	"github.com/chumnend/pook/internal/models"
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

func TestCreateBookError(t *testing.T) {
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

func TestGetAllBooks(t *testing.T) {
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

	testBooks := []models.Book{
		{
			ID:       uuid.New(),
			UserID:   uuid.New(),
			ImageURL: "https://example.com/book1.jpg",
			Title:    "Test Book 1",
		},
		{
			ID:       uuid.New(),
			UserID:   uuid.New(),
			ImageURL: "https://example.com/book2.jpg",
			Title:    "Test Book 2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "image_url", "title", "created_at", "updated_at"})
	for _, book := range testBooks {
		rows.AddRow(book.ID, book.UserID, book.ImageURL, book.Title, book.CreatedAt, book.UpdatedAt)
	}

	mock.ExpectQuery("SELECT id, user_id, image_url, title, created_at, updated_at FROM books").
		WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetAllBooks returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("GetAllBooks returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	books, ok := response["books"]
	if !ok {
		t.Errorf("Response should contain 'books' key")
	}

	booksArray, ok := books.([]interface{})
	if !ok {
		t.Errorf("Books should be an array")
	}

	if len(booksArray) != 2 {
		t.Errorf("Expected 2 books, got %d", len(booksArray))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestGetAllBooksError(t *testing.T) {
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

	mock.ExpectQuery("SELECT id, user_id, image_url, title, created_at, updated_at FROM books").
		WillReturnError(sqlmock.ErrCancelled)

	req, err := http.NewRequest("GET", "/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("GetAllBooks returned wrong status code for DB error: got %v want %v", status, http.StatusInternalServerError)
	}

	if !strings.Contains(rr.Body.String(), "unable to get all book") {
		t.Errorf("GetAllBooks should return 'unable to get all book' error message, got: %s", rr.Body.String())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestGetAllBooksByUserId(t *testing.T) {
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

	testBooks := []models.Book{
		{
			ID:       uuid.New(),
			UserID:   userID,
			ImageURL: "https://example.com/book1.jpg",
			Title:    "User Book 1",
		},
		{
			ID:       uuid.New(),
			UserID:   userID,
			ImageURL: "https://example.com/book2.jpg",
			Title:    "User Book 2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "image_url", "title", "created_at", "updated_at"})
	for _, book := range testBooks {
		rows.AddRow(book.ID, book.UserID, book.ImageURL, book.Title, book.CreatedAt, book.UpdatedAt)
	}

	mock.ExpectQuery("SELECT id, user_id, image_url, title, created_at, updated_at FROM books WHERE user_id = \\$1").
		WithArgs(userID).
		WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/v1/books?user_id="+userID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetAllBooks returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("GetAllBooks returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	books, ok := response["books"]
	if !ok {
		t.Errorf("Response should contain 'books' key")
	}

	booksArray, ok := books.([]interface{})
	if !ok {
		t.Errorf("Books should be an array")
	}

	if len(booksArray) != 2 {
		t.Errorf("Expected 2 books, got %d", len(booksArray))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestGetAllBooksByUserIdInvalidUserId(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/books?user_id=invalid-uuid", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("GetAllBooks returned wrong status code for invalid UUID: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid user id") {
		t.Errorf("GetAllBooks should return 'invalid user id' error message, got: %s", rr.Body.String())
	}
}

func TestGetAllBooksByUserIdError(t *testing.T) {
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

	mock.ExpectQuery("SELECT id, user_id, image_url, title, created_at, updated_at FROM books WHERE user_id = \\$1").
		WithArgs(userID).
		WillReturnError(sqlmock.ErrCancelled)

	req, err := http.NewRequest("GET", "/v1/books?user_id="+userID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("GetAllBooks returned wrong status code for DB error: got %v want %v", status, http.StatusInternalServerError)
	}

	if !strings.Contains(rr.Body.String(), "unable to get user's books") {
		t.Errorf("GetAllBooks should return 'unable to get user's books' error message, got: %s", rr.Body.String())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestGetBook(t *testing.T) {
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
	userID := uuid.New()

	testBook := models.Book{
		ID:       bookID,
		UserID:   userID,
		ImageURL: "https://example.com/book.jpg",
		Title:    "Test Book",
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "image_url", "title", "created_at", "updated_at"}).
		AddRow(testBook.ID, testBook.UserID, testBook.ImageURL, testBook.Title, testBook.CreatedAt, testBook.UpdatedAt)

	mock.ExpectQuery("SELECT id, user_id, image_url, title, created_at, updated_at FROM books WHERE id = \\$1").
		WithArgs(bookID).
		WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/v1/books/"+bookID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("book_id", bookID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetBook returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("GetBook returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	book, ok := response["book"]
	if !ok {
		t.Errorf("Response should contain 'book' key")
	}

	bookMap, ok := book.(map[string]interface{})
	if !ok {
		t.Errorf("Book should be an object")
	}

	if bookMap["title"] != testBook.Title {
		t.Errorf("Expected book title %s, got %v", testBook.Title, bookMap["title"])
	}

	if bookMap["imageUrl"] != testBook.ImageURL {
		t.Errorf("Expected book imageUrl %s, got %v", testBook.ImageURL, bookMap["imageUrl"])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestGetBookInvalidUserId(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/books/invalid-uuid", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("book_id", "invalid-uuid")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("GetBook returned wrong status code for invalid UUID: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid book_id") {
		t.Errorf("GetBook should return 'invalid book_id' error message, got: %s", rr.Body.String())
	}
}

func TestGetBookNotFound(t *testing.T) {
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

	mock.ExpectQuery("SELECT id, user_id, image_url, title, created_at, updated_at FROM books WHERE id = \\$1").
		WithArgs(bookID).
		WillReturnError(sqlmock.ErrCancelled)

	req, err := http.NewRequest("GET", "/v1/books/"+bookID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("book_id", bookID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("GetBook returned wrong status code for book not found: got %v want %v", status, http.StatusNotFound)
	}

	if !strings.Contains(rr.Body.String(), "book not found") {
		t.Errorf("GetBook should return 'book not found' error message, got: %s", rr.Body.String())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestUpdateBook(t *testing.T) {
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
		"imageUrl": "https://example.com/updated-book-cover.jpg",
		"title":    "Updated Test Book Title",
	}

	mock.ExpectExec("UPDATE books SET image_url = \\$1, title = \\$2, updated_at = \\$3 WHERE id = \\$4").
		WithArgs("https://example.com/updated-book-cover.jpg", "Updated Test Book Title", sqlmock.AnyArg(), bookID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("PUT", "/v1/books/"+bookID.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("book_id", bookID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("UpdateBook returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("UpdateBook returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	expectedMessage := "book successfully updated"
	if response["message"] != expectedMessage {
		t.Errorf("UpdateBook returned wrong message: got %v want %v", response["message"], expectedMessage)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestUpdateBookInvalidUserId(t *testing.T) {
	requestBody := map[string]interface{}{
		"imageUrl": "https://example.com/updated-book-cover.jpg",
		"title":    "Updated Test Book Title",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("PUT", "/v1/books/invalid-uuid", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("book_id", "invalid-uuid")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("UpdateBook returned wrong status code for invalid UUID: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "invalid book_id") {
		t.Errorf("UpdateBook should return 'invalid book_id' error message, got: %s", rr.Body.String())
	}
}

func TestUpdateBookMissingRequiredField(t *testing.T) {
	bookID := uuid.New()

	testCases := []struct {
		name string
		body map[string]interface{}
	}{
		{
			name: "missing title",
			body: map[string]interface{}{
				"imageUrl": "https://example.com/image.jpg",
			},
		},
		{
			name: "missing imageUrl",
			body: map[string]interface{}{
				"title": "Test Title",
			},
		},
		{
			name: "empty title",
			body: map[string]interface{}{
				"imageUrl": "https://example.com/image.jpg",
				"title":    "",
			},
		},
		{
			name: "empty imageUrl",
			body: map[string]interface{}{
				"imageUrl": "",
				"title":    "Test Title",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tc.body)
			req, err := http.NewRequest("PUT", "/v1/books/"+bookID.String(), bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetPathValue("book_id", bookID.String())

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(UpdateBook)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("UpdateBook returned wrong status code for %s: got %v want %v", tc.name, status, http.StatusBadRequest)
			}

			responseBody := rr.Body.String()
			if !strings.Contains(responseBody, "required") {
				t.Errorf("UpdateBook should return error message about required fields for %s, got: %s", tc.name, responseBody)
			}
		})
	}
}

func TestUpdateBookError(t *testing.T) {
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
		"imageUrl": "https://example.com/updated-book-cover.jpg",
		"title":    "Updated Test Book Title",
	}

	mock.ExpectExec("UPDATE books SET image_url = \\$1, title = \\$2, updated_at = \\$3 WHERE id = \\$4").
		WithArgs("https://example.com/updated-book-cover.jpg", "Updated Test Book Title", sqlmock.AnyArg(), bookID).
		WillReturnError(sqlmock.ErrCancelled)

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("PUT", "/v1/books/"+bookID.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("book_id", bookID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("UpdateBook returned wrong status code for DB error: got %v want %v", status, http.StatusInternalServerError)
	}

	if !strings.Contains(rr.Body.String(), "unable to update book") {
		t.Errorf("UpdateBook should return 'unable to update book' error message, got: %s", rr.Body.String())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}
