package service

import (
	"github.com/chumnend/pook/internal/entity"
)

type userService struct {
	repo entity.UserRepository
}

// NewService returns a UserService utilizing provided UserRepository
func NewService(repo entity.UserRepository) entity.UserService {
	return &userService{repo: repo}
}

func (u *userService) FindAll() ([]entity.User, error) {
	var users []entity.User
	return users, nil
}

func (u *userService) FindByUsername(email string) (*entity.User, error) {
	var user entity.User
	return &user, nil
}

func (u *userService) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	return &user, nil
}

func (u *userService) Save(user *entity.User) error {
	return nil
}

func (u *userService) Validate(user *entity.User) error {
	return nil
}

func (u *userService) GenerateToken(user *entity.User) (string, error) {
	return "", nil
}

func (u *userService) ComparePassword(user *entity.User, password string) error {
	return nil
}
