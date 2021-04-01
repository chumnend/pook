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

	"github.com/chumnend/pook/internal/pook"
	"github.com/joho/godotenv"
)

var a *pook.App

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

	a = pook.NewApp(dbURL)

	// start test runner
	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func addUser() {
	a.DB.Exec("INSERT INTO users(email, password) VALUES ($1, $2)", "tester", "test123")
}

func clearTables() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func TestApiHealthHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["message"] != "Ready to serve requests" {
		t.Errorf("Expected 'message' to be 'Ready to serve requests'. Got '%v'.", m["message"])
	}
}

func TestSpaHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	if body := res.Body.String(); !strings.Contains(body, "doctype html") {
		t.Errorf("Expected string to contain html. Got %s", body)
	}
}

func TestRegisterHandler(t *testing.T) {
	clearTables()

	var jsonStr = []byte(`{"email":"test@example.com", "password": "test123"}`)
	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["token"] == "" {
		t.Errorf("Expected 'token' to be non empty. Got %v.", m["token"])
	}
}

func TestLoginHandler(t *testing.T) {
	clearTables()

	var jsonStr = []byte(`{"email":"test@example.com", "password": "test123"}`)
	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	jsonStr = []byte(`{"email":"test@example.com", "password": "test123"}`)
	req, _ = http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res = executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["token"] == "" {
		t.Errorf("Expected 'token' to be non empty. Got %v.", m["token"])
	}
}

func TestEmptyListBooksHandler(t *testing.T) {
	clearTables()
	addUser()

	req, _ := http.NewRequest("GET", "/api/v1/books?uid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	results := m["results"].([]interface{})

	if len(results) != 0 {
		t.Errorf("Expected 'results' to be empty. Got %v.", m["results"])
	}
}
