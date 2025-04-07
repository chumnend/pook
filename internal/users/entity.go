package users

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
}

type IUserRepository interface {
	Create(*User) error
	FindAll() ([]User, error)
	FindByID(uuid.UUID) (*User, error)
	FindByUsername(string) (*User, error)
	FindByEmail(string) (*User, error)
	Update(*User) error
}

type IUserService interface {
	Create(*User) error
	FindAll() ([]User, error)
	FindByUsername(string) (*User, error)
	FindByEmail(string) (*User, error)
	Update(*User) error
	GenerateToken(*User) (string, error)
	ComparePassword(*User, string) error
}

type IUserController interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
}
