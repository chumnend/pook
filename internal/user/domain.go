package user

import (
	"context"
	"time"
)

// User strcut declaration
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"unique" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserUsecase intercase contract
type UserUsecase interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
}

// UserRepository interface contract
type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
}
