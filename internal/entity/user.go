package entity

import (
	"time"

	"github.com/gin-gonic/gin"
)

// User represents a User in the application
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	FirstName string    `gorm:"not null" json:"firstName"`
	LastName  string    `gorm:"not null" json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UserRepository is the contract between DB to the application
type UserRepository interface {
	FindAll() ([]User, error)
	FindByUsername(string) (*User, error)
	FindByEmail(string) (*User, error)
	Save(*User) error
	Migrate() error
}

// UserService handles the business logic regarding users
type UserService interface {
	FindAll() ([]User, error)
	FindByUsername(string) (*User, error)
	FindByEmail(string) (*User, error)
	Save(*User) error
	GenerateToken(*User) (string, error)
	ComparePassword(*User, string) error
}

// UserController defines user handlers in the application
type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
