package service

import (
	"errors"

	"github.com/chumnend/pook/internal/pook/entity"
	"github.com/chumnend/pook/internal/pook/repository"
)

// UserService handles the busniess logic regarding Users
type UserService interface {
	repository.UserRepository

	Validate(user *entity.User) error
	GenerateToken(user *entity.User) (string, error)
	ComparePassword(user *entity.User, password string) error
}

type userSrv struct {
	repo *repository.UserRepository
}

// NewUserService returns a UserService utilizing provided UserRepository
func NewUserService(repo *repository.UserRepository) UserService {
	return &userSrv{repo: repo}
}

func (srv *userSrv) FindAll() ([]entity.User, error) {
	var users []entity.User
	return users, errors.New("Not implemented")
}

func (srv *userSrv) FindByEmail(email string) (*entity.User, error) {
	return nil, errors.New("Not implemented")
}

func (srv *userSrv) Save(*entity.User) error {
	return errors.New("Not implemented")
}

func (srv *userSrv) Validate(user *entity.User) error {
	return errors.New("Not implemented")
}

func (srv *userSrv) GenerateToken(user *entity.User) (string, error) {
	return "", errors.New("Not implemented")
}

func (srv *userSrv) ComparePassword(user *entity.User, password string) error {
	return errors.New("Not implemented")
}
