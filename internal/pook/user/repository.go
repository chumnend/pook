package user

import (
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	conn *gorm.DB
}

// NewPostgresRepository returns a Repository struct utilizing PostgreSQL
func NewPostgresRepository(conn *gorm.DB) Repository {
	conn.AutoMigrate(&User{})
	return &userRepo{conn: conn}
}

func (repo *userRepo) FindAll() ([]User, error) {
	var users []User
	result := repo.conn.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (repo *userRepo) FindByEmail(email string) (*User, error) {
	var u User
	result := repo.conn.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return &u, result.Error
	}
	return &u, nil
}

func (repo *userRepo) Save(user *User) error {
	result := repo.conn.Create(user)
	return result.Error
}
