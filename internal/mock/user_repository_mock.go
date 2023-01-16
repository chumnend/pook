package mock

import (
	"github.com/chumnend/pook/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is mock for the UserRepository
type MockUserRepository struct {
	mock.Mock
}

// FindAll mock method
func (mock *MockUserRepository) FindAll() ([]entity.User, error) {
	args := mock.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

// FindByUsername mock method
func (mock *MockUserRepository) FindByUsername(username string) (*entity.User, error) {
	args := mock.Called(username)
	return args.Get(0).(*entity.User), args.Error(1)
}

// FindByEmail mock method
func (mock *MockUserRepository) FindByEmail(email string) (*entity.User, error) {
	args := mock.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

// Save mock method given a User struct
func (mock *MockUserRepository) Save(user *entity.User) error {
	args := mock.Called(user)
	return args.Error(0)
}

// Migrate mock method
func (mock *MockUserRepository) Migrate() error {
	args := mock.Called()
	return args.Error(0)
}
