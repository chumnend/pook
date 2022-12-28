package repository

import (
	"github.com/chumnend/pook/internal/entity"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

// NewUserPostgresRepository returns a UserRepository struct utilizing PostgreSQL
func NewUserPostgresRepository(conn *gorm.DB) entity.UserRepository {
	return &userRepository{conn: conn}
}

func (u *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User

	result := u.conn.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (u *userRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User

	result := u.conn.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	return &user, nil
}

func (u *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	result := u.conn.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	return &user, nil
}

func (u *userRepository) Save(user *entity.User) error {
	result := u.conn.Create(user)
	return result.Error
}

func (u *userRepository) Migrate() error {
	return u.conn.AutoMigrate(&entity.User{}).Error // TODO: missing tests
}
