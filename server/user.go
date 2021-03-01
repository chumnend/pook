package server

import (
	"database/sql"
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

func listUsers(db *sql.DB) ([]User, error) {
	return nil, errors.New("Not implemented")
}

func (u *User) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
