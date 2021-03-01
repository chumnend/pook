package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	bookings "github.com/chumnend/bookings-server/server"
	"github.com/joho/godotenv"
)

var s *bookings.Server

func TestMain(m *testing.M) {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file was found")
	}

	connectionString := os.Getenv("DATABASE_TEST_URL")
	if connectionString == "" {
		log.Fatal("Missing env:  DATABASE_TEST_URL")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("Missing env:  SECRET")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Missing env:  PORT")
	}

	s = bookings.NewServer(connectionString, secret, port)
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	if body := res.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestNonExistentUser(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/users/25", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, res.Code)

	var m map[string]string
	json.Unmarshal(res.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key response to be 'User not found'. Got '%v'", m["error"])
	}
}

func TestCreateUser(t *testing.T) {
	clearTable()

	var jsonStr = []byte(
		`
			{
				"email": "tester@example.com",
				"password": "test123",
				"firstName": "test",
				"lastName": "test",
			}
		`,
	)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)
	if m["email"] != "tester@example.com" {
		t.Errorf("Expected email to be 'tester@example.com'. Got '%v'", m["email"])
	}

	if m["firstName"] != "test" {
		t.Errorf("Expected firstName to be 'test'. Got '%v'", m["firstName"])
	}

	if m["lastName"] != "test" {
		t.Errorf("Expected lastName to be 'test'. Got '%v'", m["lastName"])
	}
}

func TestGetUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
}

func TestUpdateUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	res := executeRequest(req)

	var initialUser map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &initialUser)

	var jsonStr = []byte(
		`
			{
				"email": "tester@example.com",
				"password": "test123",
				"firstName": "updated",
				"lastName": "updated",
			}
		`,
	)
	req, _ = http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res = executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)
	if m["id"] != initialUser["id"] {
		t.Errorf("Expected id to be the same (%v). Got '%v'", initialUser["id"], m["id"])
	}

	if m["firstName"] == initialUser["firstName"] {
		t.Errorf("Expected firstName to change from %v. Got '%v'", initialUser["firtName"], m["firstName"])
	}

	if m["lastName"] == initialUser["lastName"] {
		t.Errorf("Expected lastName to change from %v. Got '%v'", initialUser["lastName"], m["lastName"])
	}
}

func TestDeleteUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	req, _ = http.NewRequest("DELETE", "/users/1", nil)
	res = executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	req, _ = http.NewRequest("GET", "/users/1", nil)
	res = executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
}

func clearTable() {
	s.DB.Exec("DELETE FROM users")
	s.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		countStr := strconv.Itoa(count)
		email := "tester" + countStr + "@example.com"
		password := "test" + countStr
		fname := "test" + countStr
		lname := "test" + countStr

		s.DB.Exec("INSERT INTO users(email, password, first_name, last_name) VALUES(?, ?, ?, ?)", email, password, fname, lname)
	}
}
