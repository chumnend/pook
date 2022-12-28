package service

import (
	"errors"
	"os"

	"github.com/chumnend/pook/internal/entity"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo entity.UserRepository
}

// NewUserService returns a UserService utilizing provided UserRepository
func NewUserService(repo entity.UserRepository) entity.UserService {
	return &userService{repo: repo}
}

func (u *userService) FindAll() ([]entity.User, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *userService) FindByUsername(username string) (*entity.User, error) {
	user, err := u.repo.FindByUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) FindByEmail(email string) (*entity.User, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) Save(user *entity.User) error {
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err // TODO: missing test
	}
	user.Password = string(hashedPassword)
	return u.repo.Save(user)
}

func (u *userService) GenerateToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err // TODO: missing test
	}
	return tokenStr, nil
}

func (u *userService) ComparePassword(user *entity.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
