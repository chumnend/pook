package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct declaration
type Config struct {
	DB         string
	Secret     string
	Port       string
	StaticPath string
	IndexPath  string
}

// NewConfig builds a configuration struct to be used by the application using a .env file
func NewConfig() *Config {
	var path string
	if os.Getenv("CIRCLECI") == "true" {
		path = "/home/circleci/project/"
	} else {
		path = os.ExpandEnv("$GOPATH/src/github.com/chumnend/pook/")
	}
	staticPath := path + "client/build"
	indexPath := "index.html"

	err := godotenv.Load(path + ".env")
	if err != nil {
		log.Println(".env file not found")
	}

	var databaseURL string
	if os.Getenv("ENV") != "test" {
		databaseURL = os.Getenv("DATABASE_URL")
	} else {
		databaseURL = os.Getenv("DATABASE_TEST_URL")
	}
	if databaseURL == "" {
		log.Fatal("missing env: DATABASE_URL and/or DATABASE_TEST_URL")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("missing env: SECRET_KEY")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing env: PORT")
	}

	return &Config{
		DB:         databaseURL,
		Secret:     secret,
		Port:       port,
		StaticPath: staticPath,
		IndexPath:  indexPath,
	}
}
