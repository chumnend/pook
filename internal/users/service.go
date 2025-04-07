package users

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo IUserRepository
}

func NewUserService(repo IUserRepository) IUserService {
	return &UserService{repo: repo}
}

func (srv *UserService) Create(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return srv.repo.Create(user)
}

func (srv *UserService) FindAll() ([]User, error) {
	users, err := srv.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (srv *UserService) FindByUsername(username string) (*User, error) {
	user, err := srv.repo.FindByUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (srv *UserService) FindByEmail(email string) (*User, error) {
	user, err := srv.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (srv *UserService) Update(user *User) error {
	return srv.repo.Update(user)
}

func (srv *UserService) GenerateToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (srv *UserService) ComparePassword(user *User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
