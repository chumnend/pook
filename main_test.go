package main

import (
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

var s *pook.Server

func TestMain(m *testing.M) {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	dbURL := os.Getenv("DATABASE_TEST_URL")
	if dbURL == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("missing env: SECRET_KEY")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	s = pook.NewServer(dbURL)

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

func checkIfSpa(t *testing.T, res *httptest.ResponseRecorder) {
	if body := res.Body.String(); strings.Contains(body, "doctype html") {
		t.Errorf("Expected to not hit spa handler.")
	}
}

func addTestUser() {
	s.DB.Exec("INSERT INTO users(email, password) VALUES($1, $2)", "tester@example.com", "tester")
}

func clearTable() {
	s.DB.Exec("DELETE FROM users")
	s.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func TestSpaHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); !strings.Contains(body, "doctype html") {
		t.Errorf("Expected string to contain html. Got %s", body)
	}
}

func TestGetEmptyUsers(t *testing.T) {
	clearTable()

	request, _ := http.NewRequest("GET", "/api/v1/users", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
	checkIfSpa(t, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if !m["success"].(bool) {
		t.Errorf("Expected success to be true. Got %v.", m["success"])
	}

	if len(m["payload"].([]interface{})) != 0 {
		t.Errorf("Expected empty array. Got %v.", m["payload"])
	}
}

func TestGetAllUsers(t *testing.T) {
	clearTable()
	addTestUser()

	request, _ := http.NewRequest("GET", "/api/v1/users", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
	checkIfSpa(t, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if !m["success"].(bool) {
		t.Errorf("Expected success to be true. Got %v.", m["success"])
	}

	if len(m["payload"].([]interface{})) != 1 {
		t.Errorf("Expected 1 user in array. Got %v.", m["payload"])
	}
}

func TestGetNonExistentUser(t *testing.T) {
	clearTable()

	request, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusNotFound, response.Code)
	checkIfSpa(t, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["success"].(bool) {
		t.Errorf("Expected success to be false. Got %v.", m["success"])
	}

	if m["message"] != "User not found" {
		t.Errorf("Expected 'message' to be set to 'User not found'. Got '%v'.", m["message"])
	}
}

func TestGetUser(t *testing.T) {
	clearTable()
	addTestUser()

	request, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
	checkIfSpa(t, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if !m["success"].(bool) {
		t.Errorf("Expected success to be true. Got %v.", m["success"])
	}

	if m["payload"] == "" {
		t.Errorf("Expected 'payload' to be contain user. Got %v.", m["payload"])
	}
}
