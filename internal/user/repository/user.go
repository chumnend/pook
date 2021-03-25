package repository

import (
	"context"

	"github.com/chumnend/pook/internal/user/domain"
	"github.com/jinzhu/gorm"
)

// UserRepository struct declaration
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository creats a new UserRepository
func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

// Fetch returns list of Users
func (u *UserRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := u.Conn.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetByID returns a User
func (u *UserRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	err := u.Conn.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
