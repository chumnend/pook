package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/chumnend/bookings-server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm postgres interface
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

	// migrate the schema
	db.AutoMigrate(&models.User{})

	fmt.Println("Successfully connected to db")

	return db
}
