package user

import (
	"errors"
	"os"
	"time"

	"github.com/chumnend/pook/internal/book"
	"github.com/chumnend/pook/internal/task"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User struct declaration
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"unique;not null" json:"password"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Admin     bool      `gorm:"default:false" json:"admin"`

	Books []book.Book
	Tasks []task.Task
}

// NewUser returns a new User struct
func NewUser() *User {
	return &User{}
}

// Create adds a User to the DB
func (u *User) Create(db *gorm.DB) error {
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return db.Create(&u).Error
}

// Update updates the User in the DB
func (u *User) Update(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Delete dletes the User in the DB
func (u *User) Delete(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Token struct declaration
type Token struct {
	ID    uint
	Email string
	*jwt.StandardClaims
}

// GenerateToken creates jwt token using User internal data
func (u *User) GenerateToken() (string, error) {
	tk := Token{
		ID:             u.ID,
		Email:          u.Email,
		StandardClaims: &jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (u *User) comparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}

	return true
}

// ListUsers returns a list of users in the DB
func ListUsers(db *gorm.DB) ([]User, error) {
	return nil, errors.New("Not implemented")
}

// FindUserByID takes an ID and returns a user struct
func FindUserByID(db *gorm.DB, id string) (User, error) {
	u := User{}
	return u, nil
}

// FindUserByEmail takes an ID and returns a user struct
func FindUserByEmail(db *gorm.DB, email string) (User, error) {
	var u User
	err := db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}
