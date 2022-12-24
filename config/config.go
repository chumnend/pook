package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// Config struct declaration
type Config struct {
	DB     string
	Secret string
	Port   string
}

// NewConfig builds a configuration struct to be used by the application using a .env file
func New() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	var databaseURL string
	if os.Getenv("ENV") != "test" {
		databaseURL = os.Getenv("PG_URL")
		if databaseURL == "" {
			err := errors.New("missing env: PG_URL")
			return nil, err
		}
	} else {
		databaseURL = os.Getenv("TEST_PG_URL")
		if databaseURL == "" {
			err := errors.New("missing env: TEST_PG_URL")
			return nil, err
		}
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		err := errors.New("missing env: SECRET_KEY")
		return nil, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		err := errors.New("missing env: PORT")
		return nil, err
	}

	config := &Config{
		DB:     databaseURL,
		Secret: secret,
		Port:   port,
	}

	return config, nil
}
