package service

import (
	"github.com/chumnend/pook/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock struct of a UserService
type MockUserService struct {
	mock.Mock
}

// FindAll provides mock function
func (mock *MockUserService) FindAll() ([]domain.User, error) {
	args := mock.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

// FindByEmail provides mock function
func (mock *MockUserService) FindByEmail(email string) (*domain.User, error) {
	args := mock.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}

// Save provides mock function
func (mock *MockUserService) Save(user *domain.User) error {
	args := mock.Called(user)
	return args.Error(0)
}

// Validate provides mock function
func (mock *MockUserService) Validate(user *domain.User) error {
	args := mock.Called(user)
	return args.Error(0)
}

// GenerateToken provides mock function
func (mock *MockUserService) GenerateToken(user *domain.User) (string, error) {
	args := mock.Called(user)
	return args.String(0), args.Error(1)
}

// ComparePassword provides mock function
func (mock *MockUserService) ComparePassword(user *domain.User, password string) error {
	args := mock.Called(user, password)
	return args.Error(0)
}
