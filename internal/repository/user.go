package repository

import (
	"github.com/chumnend/pook/internal/entity"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

// NewPostgresRepository returns a UserRepository struct utilizing PostgreSQL
func NewPostgresRepository(conn *gorm.DB) entity.UserRepository {
	return &userRepository{conn: conn}
}

func (u *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	return users, nil
}

func (u *userRepository) FindByUsername(email string) (*entity.User, error) {
	var user entity.User
	return &user, nil
}

func (u *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	return &user, nil
}

func (u *userRepository) Save(user *entity.User) error {
	return nil
}

func (u *userRepository) Migrate() error {
	return nil
}
