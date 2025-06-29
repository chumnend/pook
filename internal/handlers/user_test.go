package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/db"
	"github.com/chumnend/pook/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
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

	requestBody := map[string]interface{}{
		"email":    "test@example.com",
		"username": "testuser",
		"password": "password123",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(sqlmock.AnyArg(), "testuser", "test@example.com", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Register returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("Register returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	expectedMessage := "registration successful"
	if response["message"] != expectedMessage {
		t.Errorf("Register returned wrong message: got %v want %v", response["message"], expectedMessage)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestRegisterInvalidInput(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/register", strings.NewReader("invalid json"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Register returned wrong status code for invalid JSON: got %v want %v", status, http.StatusBadRequest)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	if response["message"] != "invalid input" {
		t.Errorf("Register should return 'invalid input' error message, got: %s", response["message"])
	}
}

func TestRegisterMissingRequiredFields(t *testing.T) {
	testCases := []struct {
		name string
		body map[string]interface{}
	}{
		{
			name: "missing email",
			body: map[string]interface{}{
				"username": "testuser",
				"password": "password123",
			},
		},
		{
			name: "missing username",
			body: map[string]interface{}{
				"email":    "test@example.com",
				"password": "password123",
			},
		},
		{
			name: "missing password",
			body: map[string]interface{}{
				"email":    "test@example.com",
				"username": "testuser",
			},
		},
		{
			name: "empty email",
			body: map[string]interface{}{
				"email":    "",
				"username": "testuser",
				"password": "password123",
			},
		},
		{
			name: "empty username",
			body: map[string]interface{}{
				"email":    "test@example.com",
				"username": "",
				"password": "password123",
			},
		},
		{
			name: "empty password",
			body: map[string]interface{}{
				"email":    "test@example.com",
				"username": "testuser",
				"password": "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tc.body)
			req, err := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Register)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("Register returned wrong status code for %s: got %v want %v", tc.name, status, http.StatusBadRequest)
			}

			var response map[string]string
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Errorf("Could not parse JSON response: %v", err)
			}

			if response["message"] != "all fields (email, username, password) are required" {
				t.Errorf("Register should return required fields error message for %s, got: %s", tc.name, response["message"])
			}
		})
	}
}

func TestRegisterInvalidEmail(t *testing.T) {
	requestBody := map[string]interface{}{
		"email":    "invalid-email",
		"username": "testuser",
		"password": "password123",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Register returned wrong status code for invalid email: got %v want %v", status, http.StatusBadRequest)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	if response["message"] != "invalid email format" {
		t.Errorf("Register should return 'invalid email format' error message, got: %s", response["message"])
	}
}

func TestRegisterDatabaseError(t *testing.T) {
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

	requestBody := map[string]interface{}{
		"email":    "test@example.com",
		"username": "testuser",
		"password": "password123",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(sqlmock.AnyArg(), "testuser", "test@example.com", sqlmock.AnyArg()).
		WillReturnError(sqlmock.ErrCancelled)

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Register returned wrong status code for DB error: got %v want %v", status, http.StatusInternalServerError)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestLogin(t *testing.T) {
	// Initialize config for token generation
	if config.Env == nil {
		config.Env = &config.EnvironmentVariables{
			SECRET_KEY: "test-secret-key-for-testing",
		}
	}

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
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	testUser := models.User{
		ID:        userID,
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash", "created_at"}).
		AddRow(testUser.ID, testUser.Username, testUser.Email, testUser.Password, testUser.CreatedAt)

	mock.ExpectQuery("SELECT id, username, email, password_hash, created_at FROM users WHERE username = \\$1").
		WithArgs("testuser").
		WillReturnRows(rows)

	requestBody := map[string]interface{}{
		"username": "testuser",
		"password": "password123",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Login returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("Login returned wrong content type: got %v want %v", contentType, expected)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	if response["id"] != testUser.ID.String() {
		t.Errorf("Login returned wrong user ID: got %v want %v", response["id"], testUser.ID.String())
	}

	if response["username"] != testUser.Username {
		t.Errorf("Login returned wrong username: got %v want %v", response["username"], testUser.Username)
	}

	if response["email"] != testUser.Email {
		t.Errorf("Login returned wrong email: got %v want %v", response["email"], testUser.Email)
	}

	if response["token"] == "" {
		t.Errorf("Login should return a token")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestLoginInvalidInput(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/login", strings.NewReader("invalid json"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Login returned wrong status code for invalid JSON: got %v want %v", status, http.StatusBadRequest)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	if response["message"] != "invalid input" {
		t.Errorf("Login should return 'invalid input' error message, got: %s", response["message"])
	}
}

func TestLoginMissingRequiredFields(t *testing.T) {
	testCases := []struct {
		name string
		body map[string]interface{}
	}{
		{
			name: "missing username",
			body: map[string]interface{}{
				"password": "password123",
			},
		},
		{
			name: "missing password",
			body: map[string]interface{}{
				"username": "testuser",
			},
		},
		{
			name: "empty username",
			body: map[string]interface{}{
				"username": "",
				"password": "password123",
			},
		},
		{
			name: "empty password",
			body: map[string]interface{}{
				"username": "testuser",
				"password": "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tc.body)
			req, err := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Login)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("Login returned wrong status code for %s: got %v want %v", tc.name, status, http.StatusBadRequest)
			}

			var response map[string]string
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Errorf("Could not parse JSON response: %v", err)
			}

			if response["message"] != "all fields (username, password) are required" {
				t.Errorf("Login should return required fields error message for %s, got: %s", tc.name, response["message"])
			}
		})
	}
}

func TestLoginUserNotFound(t *testing.T) {
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

	mock.ExpectQuery("SELECT id, username, email, password_hash, created_at FROM users WHERE username = \\$1").
		WithArgs("nonexistentuser").
		WillReturnError(sqlmock.ErrCancelled)

	requestBody := map[string]interface{}{
		"username": "nonexistentuser",
		"password": "password123",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Login returned wrong status code for user not found: got %v want %v", status, http.StatusBadRequest)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	if response["message"] != "invalid username and/or password" {
		t.Errorf("Login should return 'invalid username and/or password' error message, got: %s", response["message"])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestLoginInvalidPassword(t *testing.T) {
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
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

	testUser := models.User{
		ID:        userID,
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash", "created_at"}).
		AddRow(testUser.ID, testUser.Username, testUser.Email, testUser.Password, testUser.CreatedAt)

	mock.ExpectQuery("SELECT id, username, email, password_hash, created_at FROM users WHERE username = \\$1").
		WithArgs("testuser").
		WillReturnRows(rows)

	requestBody := map[string]interface{}{
		"username": "testuser",
		"password": "wrongpassword",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Login returned wrong status code for invalid password: got %v want %v", status, http.StatusBadRequest)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	if response["message"] != "invalid username and/or password" {
		t.Errorf("Login should return 'invalid username and/or password' error message, got: %s", response["message"])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestGetUser(t *testing.T) {
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
	testUser := models.User{
		ID:        userID,
		Username:  "TestUser",
		Email:     "TestUser@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash", "created_at"}).
		AddRow(testUser.ID, testUser.Username, testUser.Email, testUser.Password, testUser.CreatedAt)

	mock.ExpectQuery("SELECT id, username, email, password_hash, created_at FROM users WHERE id = \\$1").
		WithArgs(userID).
		WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/v1/users/"+userID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("user_id", userID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("GetUser returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("GetUser returned wrong content type: got %v want %v", contentType, expected)
	}

	var response models.User
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	if response.ID != testUser.ID {
		t.Errorf("Expected user ID %v, got %v", testUser.ID, response.ID)
	}

	if response.Username != testUser.Username {
		t.Errorf("Expected username %s, got %s", testUser.Username, response.Username)
	}

	if response.Email != testUser.Email {
		t.Errorf("Expected email %s, got %s", testUser.Email, response.Email)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestGetUserInvalidUserId(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/users/invalid-uuid", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("user_id", "invalid-uuid")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("GetUser returned wrong status code for invalid UUID: got %v want %v", status, http.StatusBadRequest)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	if response["message"] != "invalid user id" {
		t.Errorf("GetUser should return 'invalid user id' error message, got: %s", response["message"])
	}
}

func TestGetUserNotFound(t *testing.T) {
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

	mock.ExpectQuery("SELECT id, username, email, password_hash, created_at FROM users WHERE id = \\$1").
		WithArgs(userID).
		WillReturnError(sqlmock.ErrCancelled)

	req, err := http.NewRequest("GET", "/v1/users/"+userID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("user_id", userID.String())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("GetUser returned wrong status code for user not found: got %v want %v", status, http.StatusBadRequest)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Could not parse JSON response: %v", err)
	}

	if response["message"] != "user not found" {
		t.Errorf("GetUser should return 'user not found' error message, got: %s", response["message"])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}
