package user

import (
	"errors"
	"time"

	"github.com/chumnend/pook/internal/book"
	"github.com/chumnend/pook/internal/task"
	"github.com/jinzhu/gorm"
)

// User struct declaration
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"unique;not null" json:"id"`
	Password  string    `gorm:"unique;not null" json:"password"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Admin     bool      `gorm:"default:false" json:"admin"`

	Books []book.Book
	Tasks []task.Task
}

// NewUser returns a new User struct
func NewUser() *User {
	return &User{}
}

// Create adds a User to the DB
func (u *User) Create(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Update updates the User in the DB
func (u *User) Update(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Delete dletes the User in the DB
func (u *User) Delete(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// ListUsers returns a list of users in the DB
func ListUsers(db *gorm.DB) ([]User, error) {
	return nil, errors.New("Not implemented")
}

// FindUserByID takes an ID and returns a user struct
func FindUserByID(db *gorm.DB, id string) (User, error) {
	u := User{}
	return u, nil
}

// FindUserByEmail takes an ID and returns a user struct
func FindUserByEmail(db *gorm.DB, email string) (User, error) {
	u := User{}
	return u, nil
}
