package db

import (
	"database/sql"
	"log"

	"github.com/chumnend/pook/internal/config"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func Init() {
	var err error
	db, err := sql.Open("postgres", config.Env.PG_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("Connected to the database successfully")
}
