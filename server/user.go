package server

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// User DB Model
type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Username  string     `gorm:"unique,not null" json:"username"`
	Email     string     `gorm:"unique,not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
}

func listUsers(db *gorm.DB) error {
	return errors.New("Not Implemented")
}

func (user *User) createUser(db *gorm.DB) error {
	return errors.New("Not Implemented")
}

func (user *User) getUser(db *gorm.DB) error {
	return errors.New("Not Implemented")
}

func (user *User) updateUser(db *gorm.DB) error {
	return errors.New("Not Implemented")
}

func (user *User) deleteUser(db *gorm.DB) error {
	return errors.New("Not Implemented")
}
