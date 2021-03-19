package server

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User DB Model
type User struct {
	gorm.Model

	Username string `gorm:"unique,not null" json:"username"`
	Email    string `gorm:"unique,not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

func listUsers(db *gorm.DB) error {
	return errors.New("Not Implemented")
}

func hashPassword(password string) error {
	return errors.New("Not Implemented")
}

func checkPassword(password string) error {
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
