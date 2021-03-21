package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/chumnend/pook/server"
	"github.com/joho/godotenv"
)

var s *server.Server

func TestMain(m *testing.M) {
	// initialize app
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_TEST_URL")
	if dbURL == "" {
		log.Fatal("missing env: DATABASE_TEST_URL")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("missing env: SECRET_KEY")
	}

	s = server.New()
	s.Initialize(dbURL, port, secret)

	// start test runner
	code := m.Run()
	os.Exit(code)
}

func executeRequest(request *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, request)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func addTestUser() {
	s.DB.Exec("INSERT INTO users(username, email, password) VALUES($1, $2, $3)", "tester", "tester@example.com", "tester")
}

func clearTable() {
	s.DB.Exec("DELETE FROM users")
	s.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func TestSpaHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	res := executeRequest(request)

	checkResponseCode(t, http.StatusOK, res.Code)

	if body := res.Body.String(); !strings.Contains(body, "doctype html") {
		t.Errorf("Expected string to contain html. Got %s", body)
	}
}

func TestStatusHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/status", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "Ready to serve requests" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestListUserHandler(t *testing.T) {
	clearTable()

	request, _ := http.NewRequest("GET", "/api/v1/users", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestCreateUserHandler(t *testing.T) {
	clearTable()

	var jsonStr = []byte(`{"username":"tester", "email": "tester@example.com", "password": "tester123"}`)

	request, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	response := executeRequest(request)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected user ID to be '1'. Got '%v'", m["id"])
	}

	if m["username"] != "tester" {
		t.Errorf("Expected username to be 'tester'. Got '%v'", m["username"])
	}

	if m["email"] != "tester@example.com" {
		t.Errorf("Expected email to be 'tester@example.com'. Got '%v'", m["email"])
	}
}

func TestGetUserHandler(t *testing.T) {
	clearTable()

	request, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}

	addTestUser()

	request, _ = http.NewRequest("GET", "/api/v1/user/1", nil)
	response = executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateUserHandler(t *testing.T) {
	clearTable()
	addTestUser()

	request, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	response := executeRequest(request)

	var original map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &original)

	var jsonStr = []byte(`{"username":"tester2", "email": "tester2@example.com" }`)
	request, _ = http.NewRequest("PUT", "/api/v1/user/1", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	response = executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != original["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", original["id"], m["id"])
	}

	if m["username"] == original["username"] {
		t.Errorf("Expected the username to change from '%v' to '%v'. Got '%v'", original["username"], m["username"], m["username"])
	}

	if m["email"] == original["email"] {
		t.Errorf("Expected the email to change from '%v' to '%v'. Got '%v'", original["email"], m["email"], m["email"])
	}

	if m["password"] != original["password"] {
		t.Errorf("Expected the password to remain the same (%v). Got %v", original["password"], m["password"])
	}
}

func TestDeleteUserHandler(t *testing.T) {
	clearTable()
	addTestUser()

	request, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	request, _ = http.NewRequest("DELETE", "/api/v1/user/1", nil)
	response = executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	request, _ = http.NewRequest("GET", "/api/v1/user/1", nil)
	response = executeRequest(request)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}
