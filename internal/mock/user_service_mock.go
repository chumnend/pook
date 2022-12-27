package mock

import (
	"github.com/chumnend/pook/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock struct of a UserService
type MockUserService struct {
	mock.Mock
}

// FindAll mock method
func (mock *MockUserService) FindAll() ([]entity.User, error) {
	args := mock.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

// FindByUsername mock method
func (mock *MockUserService) FindByUsername(username string) (*entity.User, error) {
	args := mock.Called(username)
	return args.Get(0).(*entity.User), args.Error(1)
}

// FindByEmail mock method
func (mock *MockUserService) FindByEmail(email string) (*entity.User, error) {
	args := mock.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

// Save mock method
func (mock *MockUserService) Save(user *entity.User) error {
	args := mock.Called(user)
	return args.Error(0)
}

// GenerateToken mock method
func (mock *MockUserService) GenerateToken(user *entity.User) (string, error) {
	args := mock.Called(user)
	return args.String(0), args.Error(1)
}

// ComparePassword mock method
func (mock *MockUserService) ComparePassword(user *entity.User, password string) error {
	args := mock.Called(user, password)
	return args.Error(0)
}
