package domain

import "context"

// User struct declaration
type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"unique" json:"-"`
}

// UserEntity interface declaration
type UserEntity interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
}

// UserRepository interface declaration
type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
}
