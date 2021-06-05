package user

import (
	"errors"
	"os"
	"time"

	"github.com/chumnend/pook/internal/pook/board"
	"github.com/chumnend/pook/internal/pook/task"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User struct declaration
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	FirstName string    `gorm:"not null" json:"firstname"`
	LastName  string    `gorm:"not null" json:"lastname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Admin     bool      `gorm:"default:false" json:"admin"`

	Boards []board.Board `json:"boards"`
	Tasks  []task.Task   `json:"tasks"`
}

// NewUser returns a new User struct
func NewUser() *User {
	return &User{}
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

// Validate checks if user struct is correct
func (u *User) Validate() bool {
	if u.Email == "" || u.Password == "" {
		return false
	}

	return true
}

// GenerateToken creates jwt token using User internal data
func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// ComparePassword returns true if password matches password in the database
func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}

	return true
}
