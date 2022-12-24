package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chumnend/pook/config"
)

var app *App

func TestMain(m *testing.M) {
	// set test mode
	os.Setenv("ENV", "test")
	defer os.Unsetenv("ENV")

	// initialize the test application
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app = New(cfg)

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

func TestPing(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/ping", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)
	if m["message"] != "pong" {
		t.Errorf("Expected 'message' to be 'pong'; Got %v.", m["message"])
	}
}
