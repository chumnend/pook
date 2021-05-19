package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	log.SetOutput(ioutil.Discard)
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

func fillTables(numBoards int, numTasks int) {
	a.DB.Exec("INSERT INTO users(email, password, first_name, last_name) VALUES ($1, $2, $3, $4)", "tester", "123", "test_fname", "test_lname")

	for i := 0; i < numBoards; i++ {
		a.DB.Exec("INSERT INTO boards(title, user_id) VALUES ($1, $2)", "board"+strconv.Itoa(i+1), "1")

		for j := 0; j < numTasks; j++ {
			a.DB.Exec("INSERT INTO tasks(title, user_id, board_id) VALUES ($1, $2, $3)", "task"+strconv.Itoa(j+1), "1", strconv.Itoa(i+1))
		}
	}
}

func clearTables() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
	a.DB.Exec("DELETE FROM boards")
	a.DB.Exec("ALTER SEQUENCE boards_id_seq RESTART WITH 1")
	a.DB.Exec("DELETE FROM tasks")
	a.DB.Exec("ALTER SEQUENCE tasks_id_seq RESTART WITH 1")
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

func TestMissingEmailRegisterHandler(t *testing.T) {
	clearTables()

	var jsonStr = []byte(`{"password": "test123"}`)
	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["error"] != "missing and/or invalid information" {
		t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
	}
}

func TestMissingPasswordRegisterHandler(t *testing.T) {
	clearTables()

	var jsonStr = []byte(`{"email":"test@example.com"}`)
	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["error"] != "missing and/or invalid information" {
		t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
	}
}

func TestLoginHandler(t *testing.T) {
	clearTables()

	var jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
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

func TestBadEmailLoginHandler(t *testing.T) {
	clearTables()

	jsonStr := []byte(`{"email":"test@example.com", "password": "123"}`)
	req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["error"] != "invalid email and/or password" {
		t.Errorf("Expected the 'error' to be 'invalid email and/or password'. Got '%v'", m["error"])
	}
}

func TestBadPasswordLoginHandler(t *testing.T) {
	clearTables()

	var jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	jsonStr = []byte(`{"email":"test@example.com", "password": "567"}`)
	req, _ = http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res = executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["error"] != "invalid email and/or password" {
		t.Errorf("Expected the 'error' to be 'invalid email and/or password'. Got '%v'", m["error"])
	}
}

func TestEmptyListBoardsHandler(t *testing.T) {
	clearTables()
	fillTables(0, 0)

	req, _ := http.NewRequest("GET", "/api/v1/boards?userid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["results"]; !ok {
		t.Errorf("Expected `results` to exist. Got '%v'", m)
		return
	}

	results := m["results"].([]interface{})

	if len(results) != 0 {
		t.Errorf("Expected 'results' to be empty. Got %v.", m["results"])
	}
}

func TestListBoardsHandler(t *testing.T) {
	clearTables()
	fillTables(3, 0)

	req, _ := http.NewRequest("GET", "/api/v1/boards?userid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["results"]; !ok {
		t.Errorf("Expected `results` to exist. Got '%v'", m)
		return
	}

	results := m["results"].([]interface{})

	if len(results) != 3 {
		t.Errorf("Expected 'results' to have length of 3. Got %v.", m["results"])
	}
}

func TestCreateBoard(t *testing.T) {
	clearTables()

	var jsonStr = []byte(`{"title":"test"}`)
	req, _ := http.NewRequest("POST", "/api/v1/boards?userid=1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["result"]; !ok {
		t.Errorf("Expected `result` to exist. Got '%v'", m)
		return
	}

	result := m["result"].(map[string]interface{})

	if result["id"] != 1.0 {
		t.Errorf("Expected `id` to be '1'. Got '%v'", m["id"])
	}

	if result["title"] != "test" {
		t.Errorf("Expected 'title' to be 'test'. Got '%v'", m["title"])
	}
}

func TestNonExistentGetBoard(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/api/v1/board/1?userid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["error"] != "board not found" {
		t.Errorf("Expected the 'error' to be 'board not found'. Got '%v'", m["error"])
	}
}

func TestGetBoard(t *testing.T) {
	clearTables()
	fillTables(1, 0)

	req, _ := http.NewRequest("GET", "/api/v1/board/1?userid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["result"]; !ok {
		t.Errorf("Expected `result` to exist. Got '%v'", m)
		return
	}

	result := m["result"].(map[string]interface{})

	if result["id"] != 1.0 {
		t.Errorf("Expected `id` to be '1'. Got '%v'", m["id"])
	}

	if result["title"] != "board1" {
		t.Errorf("Expected 'title' to be 'board1'. Got '%v'", m["title"])
	}
}

func TestUpdateBoard(t *testing.T) {
	clearTables()
	fillTables(1, 0)

	req, _ := http.NewRequest("GET", "/api/v1/board/1?userid=1", nil)
	res := executeRequest(req)
	var original map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &original)

	var jsonStr = []byte(`{"title": "new title", "body": "new body"}`)
	req, _ = http.NewRequest("PUT", "/api/v1/board/1?userid=1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res = executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["result"]; !ok {
		t.Errorf("Expected `result` to exist. Got '%v'", m)
		return
	}

	orig := original["result"].(map[string]interface{})
	result := m["result"].(map[string]interface{})

	if result["id"] != orig["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", orig["id"], result["id"])
	}

	if result["title"] == orig["title"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", orig["title"], result["title"], result["title"])
	}

	if result["body"] == orig["body"] {
		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", orig["body"], result["body"], result["body"])
	}
}

func TestDeleteBoard(t *testing.T) {
	clearTables()
	fillTables(1, 0)

	req, _ := http.NewRequest("GET", "/api/v1/board/1?userid=1", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	req, _ = http.NewRequest("DELETE", "/api/v1/board/1?userid=1", nil)
	res = executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	req, _ = http.NewRequest("GET", "/api/v1/board/1?userid=1", nil)
	res = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, res.Code)
}

func TestEmptyListTasksHandler(t *testing.T) {
	clearTables()
	fillTables(1, 0)

	req, _ := http.NewRequest("GET", "/api/v1/tasks?userid=1&boardid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["results"]; !ok {
		t.Errorf("Expected `results` to exist. Got '%v'", m)
		return
	}

	results := m["results"].([]interface{})

	if len(results) != 0 {
		t.Errorf("Expected 'results' to be empty. Got %v.", m["results"])
	}
}

func TestListTasksHandler(t *testing.T) {
	clearTables()
	fillTables(1, 3)

	req, _ := http.NewRequest("GET", "/api/v1/tasks?userid=1&boardid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["results"]; !ok {
		t.Errorf("Expected `results` to exist. Got '%v'", m)
		return
	}

	results := m["results"].([]interface{})

	if len(results) != 3 {
		t.Errorf("Expected 'results' to have length of 3. Got %v.", m["results"])
	}
}

func TestCreateTaskHandler(t *testing.T) {
	clearTables()
	fillTables(1, 0)

	var jsonStr = []byte(`{"title":"test", "body":"test"}`)
	req, _ := http.NewRequest("POST", "/api/v1/tasks?userid=1&boardid=1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["result"]; !ok {
		t.Errorf("Expected `result` to exist. Got '%v'", m)
		return
	}

	result := m["result"].(map[string]interface{})

	if result["id"] != 1.0 {
		t.Errorf("Expected `id` to be '1'. Got '%v'", m["id"])
	}

	if result["title"] != "test" {
		t.Errorf("Expected 'title' to be 'test'. Got '%v'", m["title"])
	}

	if result["body"] != "test" {
		t.Errorf("Expected 'body' to be 'test'. Got '%v'", m["body"])
	}
}

func TestNonExistentGetTaskHandler(t *testing.T) {
	clearTables()
	fillTables(1, 0)

	req, _ := http.NewRequest("GET", "/api/v1/task/1?userid=1&boardid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["error"] != "task not found" {
		t.Errorf("Expected the 'error' to be 'task not found'. Got '%v'", m["error"])
	}
}

func TestGetTaskHandler(t *testing.T) {
	clearTables()
	fillTables(1, 1)

	req, _ := http.NewRequest("GET", "/api/v1/task/1?userid=1&boardid=1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["result"]; !ok {
		t.Errorf("Expected `result` to exist. Got '%v'", m)
		return
	}

	result := m["result"].(map[string]interface{})

	if result["id"] != 1.0 {
		t.Errorf("Expected `id` to be '1'. Got '%v'", m["id"])
	}

	if result["title"] != "task1" {
		t.Errorf("Expected 'title' to be 'task1'. Got '%v'", m["title"])
	}
}

func TestUpdateTaskHandler(t *testing.T) {
	clearTables()
	fillTables(1, 1)

	req, _ := http.NewRequest("GET", "/api/v1/task/1?userid=1&boardid=1", nil)
	res := executeRequest(req)
	var original map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &original)

	var jsonStr = []byte(`{"title": "new title", "body": "new body"}`)
	req, _ = http.NewRequest("PUT", "/api/v1/task/1?userid=1&boardid=1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res = executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if _, ok := m["result"]; !ok {
		t.Errorf("Expected `result` to exist. Got '%v'", m)
		return
	}

	orig := original["result"].(map[string]interface{})
	result := m["result"].(map[string]interface{})

	if result["id"] != orig["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", orig["id"], result["id"])
	}

	if result["title"] == orig["title"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", orig["title"], result["title"], result["title"])
	}

	if result["body"] == orig["body"] {
		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", orig["body"], result["body"], result["body"])
	}
}

func TestDeleteTaskHandler(t *testing.T) {
	clearTables()
	fillTables(1, 1)

	req, _ := http.NewRequest("GET", "/api/v1/task/1?userid=1&boardid=1", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	req, _ = http.NewRequest("DELETE", "/api/v1/task/1?userid=1&boardid=1", nil)
	res = executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	req, _ = http.NewRequest("GET", "/api/v1/task/1?userid=1&boardid=1", nil)
	res = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, res.Code)
}
