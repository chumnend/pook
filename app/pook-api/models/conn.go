package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connect(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	DB = db

	return db, nil
}
