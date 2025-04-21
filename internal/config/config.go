package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariables struct {
	PORT       string
	PG_URL     string
	SECRET_KEY string
}

var Env *EnvironmentVariables

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	pgUrl := os.Getenv("PG_URL")
	if pgUrl == "" {
		panic("PG_URL environment variable is not set")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		panic("SECRET_KEY environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT environment variable is not set")
	}

	Env = &EnvironmentVariables{
		PORT:       port,
		PG_URL:     pgUrl,
		SECRET_KEY: secretKey,
	}
}
