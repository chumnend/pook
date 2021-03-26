package user

import (
	"context"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

// NewUserRepository creats a new UserRepository
func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepository{
		Conn: conn,
	}
}

// Fetch returns list of Users
func (u *userRepository) Fetch(ctx context.Context) ([]User, error) {
	var users []User
	err := u.Conn.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetByID returns a User
func (u *userRepository) GetByID(ctx context.Context, id string) (User, error) {
	var user User
	err := u.Conn.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetByEmail returns a User
func (u *userRepository) GetByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := u.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
