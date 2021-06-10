package user

import (
	"github.com/chumnend/pook/internal/pook/domain"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	conn *gorm.DB
}

// NewPostgresRepository returns a UserRepository struct utilizing PostgreSQL
func NewPostgresRepository(conn *gorm.DB) domain.UserRepository {
	conn.AutoMigrate(&domain.User{})
	return &userRepo{conn: conn}
}

func (repo *userRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	result := repo.conn.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (repo *userRepo) FindByEmail(email string) (*domain.User, error) {
	var u domain.User
	result := repo.conn.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return &u, result.Error
	}
	return &u, nil
}

func (repo *userRepo) Save(user *domain.User) error {
	result := repo.conn.Create(user)
	return result.Error
}
