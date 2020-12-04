package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//User struct declaration
type User struct {
	gorm.Model

	Email     string `gorm:"type:varchar(100);unique;not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	FirstName string `gorm:"type:varchar(100)"`
	LastName  string `gorm:"type:varchar(100)"`
	Gender    string `gorm:"type:varchar(1)"`
	BirthDate time.Time
}
