package postgres

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

// Conn struct declaration
type Conn struct {
	DB *gorm.DB
}

// NewConnection creates a new connection to postgresql DB
func NewConnection(url string) *Conn {
	db, err := gorm.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	return &Conn{DB: db}
}
