package users

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the User table
type User struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
}

// UserRepository is an interface for accessing the users table is PostgreSQL
type IUserRepository interface {
	Create(*User) error
	FindAll() ([]User, error)
	FindByID(uuid.UUID) (*User, error)
	FindByUsername(string) (*User, error)
	FindByEmail(string) (*User, error)
	Update(*User) error
}
