package entity

import (
	"context"

	"github.com/chumnend/pook/internal/user/domain"
)

// UserEntity struct declaration
type UserEntity struct {
	Repo domain.UserRepository
}

// NewUserEntity creates new UserEntity
func NewUserEntity(repo domain.UserRepository) *UserEntity {
	return &UserEntity{
		Repo: repo,
	}
}

// Fetch returns a list of Users
func (u *UserEntity) Fetch(ctx context.Context) ([]domain.User, error) {
	users, err := u.Repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetByID returns a User
func (u *UserEntity) GetByID(ctx context.Context, id string) (domain.User, error) {
	user, err := u.Repo.GetByID(ctx, id)
	if err != nil {
		return user, err
	}

	return user, nil
}
