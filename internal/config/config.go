package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct declaration
type Config struct {
	Env        string
	StaticPath string
	IndexPath  string

	DB     string
	Secret string
	Port   string
}

// GetEnv returns the configuation for the application by reading the .env file
func GetEnv() *Config {
	config := Config{}

	config.Env = "prod"
	config.StaticPath = os.ExpandEnv("$GOPATH/src/github.com/chumnend/pook/web/build")
	config.IndexPath = "index.html"

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	config.DB = os.Getenv("DATABASE_URL")
	if config.DB == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	config.Secret = os.Getenv("SECRET_KEY")
	if config.Secret == "" {
		log.Fatal("missing env: SECRET_KEY")
	}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		log.Fatal("missing env: PORT")
	}

	return &config
}

// GetTestEnv returns the configuration for the testing by reading the .env file
func GetTestEnv() *Config {
	config := Config{}

	config.Env = "test"
	config.StaticPath = os.ExpandEnv("$GOPATH/src/github.com/chumnend/pook/web/build")
	config.IndexPath = "index.html"

	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/chumnend/pook/.env")) // fix to run .env not in root
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.DB = os.Getenv("DATABASE_TEST_URL")
	if config.DB == "" {
		log.Fatal("missing env: DATABASE_TEST_URL")
	}

	config.Secret = os.Getenv("SECRET_KEY")
	if config.Secret == "" {
		log.Fatal("missing env: SECRET_KEY")
	}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		log.Fatal("missing env: PORT")
	}

	return &config
}
