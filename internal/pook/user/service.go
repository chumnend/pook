package user

import (
	"errors"
)

// Service handles the busniess logic regarding Users
type Service interface {
	Repository
	Validate(user *User) error
	GenerateToken(user *User) (string, error)
	ComparePassword(user *User, password string) error
}

type userSrv struct {
	repo Repository
}

// NewService returns a Service utilizing provided Repository
func NewService(repo Repository) Service {
	return &userSrv{repo: repo}
}

func (srv *userSrv) FindAll() ([]User, error) {
	var users []User
	return users, errors.New("Not implemented")
}

func (srv *userSrv) FindByEmail(email string) (*User, error) {
	return nil, errors.New("Not implemented")
}

func (srv *userSrv) Save(*User) error {
	return errors.New("Not implemented")
}

func (srv *userSrv) Validate(user *User) error {
	return errors.New("Not implemented")
}

func (srv *userSrv) GenerateToken(user *User) (string, error) {
	return "", errors.New("Not implemented")
}

func (srv *userSrv) ComparePassword(user *User, password string) error {
	return errors.New("Not implemented")
}
