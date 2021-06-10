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

// GetEnv returns the configuation for the application by reading the .env file
func GetEnv() *Config {
	path := os.ExpandEnv("$GOPATH/src/github.com/chumnend/pook/")

	err := godotenv.Load(path + ".env")
	if err != nil {
		log.Fatal(".env file not found")
	}

	var databaseURL string
	if os.Getenv("ENV") != "test" {
		databaseURL = os.Getenv("DATABASE_URL")
	} else {
		databaseURL = os.Getenv("DATABASE_TEST_URL")
	}
	if databaseURL == "" {
		log.Fatal("missing env: DATABASE_URL")
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
		StaticPath: path + "web/build",
		IndexPath:  path + "index.html",
	}
}
