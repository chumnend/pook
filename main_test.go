package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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

	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		log.Fatal("Missing env:  DATABASE_URL")
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
