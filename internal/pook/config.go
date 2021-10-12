package pook

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

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
		_, b, _, _ := runtime.Caller(0)
		root := filepath.Join(filepath.Dir(b), "../../")
		path = os.ExpandEnv(root)
	}
	staticPath := path + "react/build"
	indexPath := "index.html"

	err := godotenv.Load(path + ".env")
	if err != nil {
		log.Println(".env file not found")
		log.Println(err)
	}

	var databaseURL string
	if os.Getenv("ENV") != "test" {
		databaseURL = os.Getenv("DATABASE_URL")
	} else {
		databaseURL = os.Getenv("TEST_DATABASE_URL")
	}
	if databaseURL == "" {
		log.Fatal("missing env: DATABASE_URL and/or TEST_DATABASE_URL")
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
