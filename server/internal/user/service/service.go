package service

import (
	"errors"
	"os"

	"github.com/chumnend/pook/server/internal/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userSrv struct {
	repo domain.UserRepository
}

// NewService returns a UserService utilizing provided UserRepository
func NewService(repo domain.UserRepository) domain.UserService {
	return &userSrv{repo: repo}
}

func (srv *userSrv) FindAll() ([]domain.User, error) {
	users, err := srv.repo.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (srv *userSrv) FindByEmail(email string) (*domain.User, error) {
	user, err := srv.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (srv *userSrv) Save(user *domain.User) error {
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return srv.repo.Save(user)
}

func (srv *userSrv) Validate(user *domain.User) error {
	if user == nil {
		return errors.New("user is empty")
	}
	if user.Email == "" || user.Password == "" {
		return errors.New("invalid user")
	}
	return nil
}

func (srv *userSrv) GenerateToken(user *domain.User) (string, error) {
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

func (srv *userSrv) ComparePassword(user *domain.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
