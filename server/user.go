package server

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User struct declaration
type User struct {
	gorm.Model

	Email     string `gorm:"type:varchar(50);unique;not null" json:"email"`
	Password  string `gorm:"type:varchar(50);not null" json:"password"`
	FirstName string `gorm:"type:varchar(50)" json:"firstName"`
	LastName  string `gorm:"type:varchar(50)" json:"lastName"`
}

func (u *User) ListUsers(db *gorm.DB) ([]User, error) {
	return nil, errors.New("Not implemented")
}

func (u *User) GetUser(db *gorm.DB) ([]User, error) {
	return nil, errors.New("Not implemented")
}

func (u *User) CreateUser(db *gorm.DB) ([]User, error) {
	return nil, errors.New("Not implemented")
}

func (u *User) UpdateUser(db *gorm.DB) ([]User, error) {
	return nil, errors.New("Not implemented")
}

func (u *User) DeleteUser(db *gorm.DB) ([]User, error) {
	return nil, errors.New("Not implemented")
}

func Register(db *gorm.DB) error {
	return errors.New("Not implemented")
}

func Login(db *gorm.DB) error {
	return errors.New("Not implemented")
}
