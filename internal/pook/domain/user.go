package domain

import (
	"net/http"
	"time"
)

// User represents a User in the application
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	FirstName string    `gorm:"not null" json:"firstName"`
	LastName  string    `gorm:"not null" json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsAdmin   bool      `gorm:"default:false" json:"isAdmin"`

	Books []Book `json:"books"`
}

// UserRepository is the contract between DB to the application
type UserRepository interface {
	FindAll() ([]User, error)
	FindByEmail(string) (*User, error)
	Save(*User) error
	Migrate() error
}

// UserService handles the business logic regarding Users
type UserService interface {
	FindAll() ([]User, error)
	FindByEmail(string) (*User, error)
	Save(*User) error
	Validate(*User) error
	GenerateToken(*User) (string, error)
	ComparePassword(*User, string) error
}

// UserController defines user handlers in the application
type UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
