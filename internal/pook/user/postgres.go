package user

import (
	"errors"

	"github.com/chumnend/pook/internal/pook/postgres"
)

type userRepo struct {
	conn *postgres.Conn
}

// NewPostgresRepository returns a Repository struct utilizing PostgreSQL
func NewPostgresRepository(conn *postgres.Conn) Repository {
	conn.DB.AutoMigrate(&User{})
	return &userRepo{conn: conn}
}

func (repo *userRepo) FindAll() ([]User, error) {
	var users []User
	return users, errors.New("Not implemented")
}

func (repo *userRepo) FindByEmail(email string) (*User, error) {
	return nil, errors.New("Not implemented")
}

func (repo *userRepo) Save(*User) error {
	return errors.New("Not implemented")
}
