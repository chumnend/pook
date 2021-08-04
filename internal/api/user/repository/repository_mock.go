package repository

import (
	"github.com/chumnend/pook/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is mock for the UserRepository
type MockUserRepository struct {
	mock.Mock
}

// FindAll mock method
func (mock *MockUserRepository) FindAll() ([]domain.User, error) {
	args := mock.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

// FindByEmail mock method
func (mock *MockUserRepository) FindByEmail(email string) (*domain.User, error) {
	args := mock.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}

// Save mock method given a User struct
func (mock *MockUserRepository) Save(user *domain.User) error {
	args := mock.Called(user)
	return args.Error(0)
}

// Migrate mock method
func (mock *MockUserRepository) Migrate() error {
	args := mock.Called()
	return args.Error(0)
}
