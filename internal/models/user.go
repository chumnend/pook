package models

import (
	"errors"
	"time"

	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password_hash"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

func CreateUser(username string, email string, password string) error {
	uuid := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("INSERT INTO users (id, username, email, password_hash) VALUES ($1, $2, $3, $4)", uuid, username, email, hashedPassword)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return errors.New("user with the same email or username already exists")
			default:
				return errors.New("unable to create user, please try again later")
			}
		}
		return err
	}

	return nil
}

func GetUserByID(id uuid.UUID) (*User, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, username, email, password_hash, created_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, username, email, password_hash, created_at FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, username, email, password_hash, created_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GenerateUserToken(user *User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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
