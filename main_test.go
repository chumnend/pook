package main

import (
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

	s = pook.NewServer()

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

func TestSpaHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	res := executeRequest(request)

	checkResponseCode(t, http.StatusOK, res.Code)

	if body := res.Body.String(); !strings.Contains(body, "doctype html") {
		t.Errorf("Expected string to contain html. Got %s", body)
	}
}
