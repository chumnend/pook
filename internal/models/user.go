package models

import (
	"errors"
	"time"

	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

func CreateUser(username string, email string, password string) error {
	uuid := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)", uuid, username, email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(id uuid.UUID) (*User, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, username, email, password, created_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, username, email, password, created_at FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, username, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GenerateUserToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})
	tokenStr, err := token.SignedString([]byte(config.Env.SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ComparePassword(user *User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
