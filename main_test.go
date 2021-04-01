package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
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

func addUser(numBooks int) {
	a.DB.Exec("INSERT INTO users(email, password) VALUES ($1, $2)", "tester", "test123")

	for i := 0; i < numBooks; i++ {
		a.DB.Exec("INSERT INTO books(title, user_id) VALUES ($1, $2)", "test"+strconv.Itoa(i+1), "1")
	}
}

func clearTables() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
	a.DB.Exec("DELETE FROM books")
	a.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
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
	addUser(0)

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

func TestListBooksHandler(t *testing.T) {
	clearTables()
	addUser(3)

	req, _ := http.NewRequest("GET", "/api/v1/books?uid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	results := m["results"].([]interface{})

	if len(results) != 3 {
		t.Errorf("Expected 'results' to have length of 3. Got %v.", m["results"])
	}
}

func TestCreateBook(t *testing.T) {
	clearTables()

	var jsonStr = []byte(`{"title":"test"}`)
	req, _ := http.NewRequest("POST", "/api/v1/books?uid=1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	result := m["result"].(map[string]interface{})

	if result["id"] != 1.0 {
		t.Errorf("Expected `id` to be '1'. Got '%v'", m["id"])
	}

	if result["title"] != "test" {
		t.Errorf("Expected 'title' to be 'test'. Got '%v'", m["title"])
	}
}

func TestNonExistentGetBook(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/api/v1/book/1?uid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["error"] != "book not found" {
		t.Errorf("Expected the 'error' to be 'book not found'. Got '%v'", m["error"])
	}
}

func TestGetBook(t *testing.T) {
	clearTables()
	addUser(1)

	req, _ := http.NewRequest("GET", "/api/v1/book/1?uid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	result := m["result"].(map[string]interface{})

	if result["id"] != 1.0 {
		t.Errorf("Expected `id` to be '1'. Got '%v'", m["id"])
	}

	if result["title"] != "test1" {
		t.Errorf("Expected 'title' to be 'test1'. Got '%v'", m["title"])
	}
}

func TestUpdateBook(t *testing.T) {
	clearTables()
	t.Errorf("Test not implemented")
}

func TestDeleteBook(t *testing.T) {
	clearTables()
	t.Errorf("Test not implemented")
}
