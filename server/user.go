package server

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User DB Model
type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Username  string     `gorm:"unique,not null" json:"username"`
	Email     string     `gorm:"unique,not null" json:"email"`
	Password  string     `gorm:"not null" json:"-"`
}

func listUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (user *User) createUser(db *gorm.DB) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	result := db.Create(&user)
	return result.Error
}

func (user *User) getUser(db *gorm.DB) error {
	result := db.First(&user)
	return result.Error
}

func (user *User) updateUser(db *gorm.DB) error {
	result := db.Save(&user)
	return result.Error
}

func (user *User) deleteUser(db *gorm.DB) error {
	result := db.Delete(&user)
	return result.Error
}
