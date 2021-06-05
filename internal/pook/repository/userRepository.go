package repository

import (
	"github.com/chumnend/pook/internal/pook/entity"
)

// UserRepository is the contract between DB to the application
type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
}
