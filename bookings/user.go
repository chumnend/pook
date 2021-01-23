package bookings

import (
	"github.com/jinzhu/gorm"
)

// User struct declaration
type User struct {
	gorm.Model

	Email     string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string `gorm:"type:varchar(100);not null" json:"-"`
	FirstName string `gorm:"type:varchar(100)" json:"firstName"`
	LastName  string `gorm:"type:varchar(100)" json:"lastName"`
}
