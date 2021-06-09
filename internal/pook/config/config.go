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

// LoadEnv returns the configuation for the application by reading the .env file
func LoadEnv() *Config {
	envPath := os.ExpandEnv("$GOPATH/src/github.com/chumnend/pook/")

	err := godotenv.Load(envPath + ".env")
	if err != nil {
		log.Fatal(".env file not found")
	}

	databaseURL := os.Getenv("DATABASE_URL")
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
		StaticPath: "web/build",
		IndexPath:  "index.html",
	}
}
