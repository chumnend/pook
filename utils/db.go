package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// ConnectDB - make database connections
func ConnectDB() *gorm.DB {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")

	// connect using database url
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}

	fmt.Println("Successfully connected to db", db)

	return db
}
