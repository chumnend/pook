package pook_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/chumnend/pook/server/internal/pook"
)

var app *pook.App

func TestMain(m *testing.M) {
	// set test mode
	os.Setenv("ENV", "test")
	defer os.Unsetenv("ENV")

	// initialize the test application
	app = pook.NewApp()

	// start test runner
	log.SetOutput(ioutil.Discard)
	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func emptyDB() {
	app.DB.Exec("DELETE FROM users")
	app.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func TestSpaHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	if body := res.Body.String(); !strings.Contains(body, "doctype html") {
		t.Errorf("Expected string to contain html. Got %s", body)
	}
}

func TestRegister(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("success", func(t *testing.T) {
		emptyDB()

		var jsonStr = []byte(`{"email":"test@example.com", "password": "test123"}`)
		req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		res := executeRequest(req)

		checkResponseCode(t, http.StatusOK, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["token"] == "" {
			t.Errorf("Expected 'token' to be non empty. Got %v.", m["token"])
		}
	})

	t.Run("fail - no email", func(t *testing.T) {
		emptyDB()

		var jsonStr = []byte(`{"password": "test123"}`)
		req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		res := executeRequest(req)

		checkResponseCode(t, http.StatusBadRequest, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["error"] != "missing and/or invalid information" {
			t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - no password", func(t *testing.T) {
		emptyDB()

		var jsonStr = []byte(`{"email":"test@example.com"}`)
		req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		res := executeRequest(req)

		checkResponseCode(t, http.StatusBadRequest, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["error"] != "missing and/or invalid information" {
			t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
		}
	})
}

func TestLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("success", func(t *testing.T) {
		emptyDB()

		var jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
		req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		res := executeRequest(req)

		checkResponseCode(t, http.StatusOK, res.Code)

		jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
		req, _ = http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		res = executeRequest(req)

		checkResponseCode(t, http.StatusOK, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["token"] == "" {
			t.Errorf("Expected 'token' to be non empty. Got %v.", m["token"])
		}
	})

	t.Run("fail - bad email", func(t *testing.T) {
		emptyDB()

		jsonStr := []byte(`{"email":"test@example.com", "password": "123"}`)
		req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		res := executeRequest(req)

		checkResponseCode(t, http.StatusBadRequest, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["error"] != "invalid email and/or password" {
			t.Errorf("Expected the 'error' to be 'invalid email and/or password'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - bad password", func(t *testing.T) {
		emptyDB()

		var jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
		req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		res := executeRequest(req)

		checkResponseCode(t, http.StatusOK, res.Code)

		jsonStr = []byte(`{"email":"test@example.com", "password": "567"}`)
		req, _ = http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		res = executeRequest(req)

		checkResponseCode(t, http.StatusBadRequest, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["error"] != "invalid email and/or password" {
			t.Errorf("Expected the 'error' to be 'invalid email and/or password'. Got '%v'", m["error"])
		}
	})
}
