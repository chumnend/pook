package server

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User struct declaration
type User struct {
	gorm.Model

	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Bookings  []Booking
	Locations []Location
}

func listUsers(db *gorm.DB) ([]User, error) {
	return nil, errors.New("Not implemented")
}

func (u *User) createUser(db *gorm.DB) error {
	return errors.New("Not implemented")
}

func (u *User) getUser(db *gorm.DB) error {
	return errors.New("Not implemented")
}

func (u *User) updateUser(db *gorm.DB) error {
	return errors.New("Not implemented")
}

func (u *User) deleteUser(db *gorm.DB) error {
	return errors.New("Not implemented")
}
