package repository

import (
	"errors"

	"github.com/chumnend/pook/internal/pook/entity"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	conn *gorm.DB
}

// NewUserRepository returns a UseeRepository struct utilizing PostgreSQL
func NewUserRepository(conn *gorm.DB) UserRepository {
	conn.AutoMigrate(&entity.User{})
	return &userRepo{conn: conn}
}

func (repo *userRepo) FindAll() ([]entity.User, error) {
	var users []entity.User
	return users, errors.New("Not implemented")
}

func (repo *userRepo) FindByEmail(email string) (*entity.User, error) {
	return nil, errors.New("Not implemented")
}

func (repo *userRepo) Save(*entity.User) error {
	return errors.New("Not implemented")
}
