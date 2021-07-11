package repository

import (
	"github.com/chumnend/pook/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is mock for the UserRepository
type MockUserRepository struct {
	mock.Mock
}

// FindAll provides a mock function
func (mock *MockUserRepository) FindAll() ([]domain.User, error) {
	args := mock.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

// FindByEmail provides a mock function
func (mock *MockUserRepository) FindByEmail(email string) (*domain.User, error) {
	args := mock.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}

// Save provides a mock function given a User struct
func (mock *MockUserRepository) Save(user *domain.User) error {
	args := mock.Called(user)
	return args.Error(0)
}

// Migrate provides a mock function
func (mock *MockUserRepository) Migrate() error {
	args := mock.Called()
	return args.Error(0)
}
