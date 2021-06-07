package user

import (
	"github.com/chumnend/pook/internal/pook/postgres"
)

type userRepo struct {
	conn *postgres.Conn
}

// NewPostgresRepository returns a Repository struct utilizing PostgreSQL
func NewPostgresRepository(conn *postgres.Conn) Repository {
	conn.DB.AutoMigrate(&User{})
	return &userRepo{conn: conn}
}

func (repo *userRepo) FindAll() ([]User, error) {
	var users []User
	result := repo.conn.DB.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (repo *userRepo) FindByEmail(email string) (*User, error) {
	var u User
	result := repo.conn.DB.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return &u, result.Error
	}
	return &u, nil
}

func (repo *userRepo) Save(user *User) error {
	result := repo.conn.DB.Create(user)
	return result.Error
}
