package postgres

import (
	"github.com/chumnend/pook/internal/pook/user"
)

type userRepo struct {
	conn *Conn
}

// NewUserRepository returns a Repository struct utilizing PostgreSQL
func NewUserRepository(conn *Conn) user.Repository {
	conn.DB.AutoMigrate(&user.User{})
	return &userRepo{conn: conn}
}

func (repo *userRepo) FindAll() ([]user.User, error) {
	var users []user.User
	result := repo.conn.DB.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (repo *userRepo) FindByEmail(email string) (*user.User, error) {
	var u user.User
	result := repo.conn.DB.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return &u, result.Error
	}
	return &u, nil
}

func (repo *userRepo) Save(user *user.User) error {
	result := repo.conn.DB.Create(user)
	return result.Error
}
