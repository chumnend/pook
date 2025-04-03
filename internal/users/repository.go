package users

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	return nil
}

func (r *UserRepository) FindAll() ([]User, error) {
	var users []User

	return users, nil
}

func (r *UserRepository) FindByID(id uuid.UUID) (*User, error) {
	return &User{}, nil
}

func (r *UserRepository) FindByUsername(username string) (*User, error) {
	return &User{}, nil
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	return &User{}, nil
}

func (r *UserRepository) Update(user *User) error {
	return nil
}
