package mock

import (
	"github.com/chumnend/pook/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockPageService is a mock struct of a BookService
type MockPageService struct {
	mock.Mock
}

// FindAllByBookID mock method
func (mock *MockPageService) FindAllByBookID(id uint) ([]entity.Page, error) {
	args := mock.Called(id)
	return args.Get(0).([]entity.Page), args.Error(1)
}

// FindByID mock method
func (mock *MockPageService) FindByID(id uint) (*entity.Page, error) {
	args := mock.Called(id)
	return args.Get(0).(*entity.Page), args.Error(1)
}

// Create mock method
func (mock *MockPageService) Create(page *entity.Page) error {
	args := mock.Called(page)
	return args.Error(0)
}

// Update mock method
func (mock *MockPageService) Update(page *entity.Page) error {
	args := mock.Called(page)
	return args.Error(0)
}

// Delete mock method
func (mock *MockPageService) Delete(page *entity.Page) error {
	args := mock.Called(page)
	return args.Error(0)
}

// Validate mock method
func (mock *MockPageService) Validate(page *entity.Page) error {
	args := mock.Called(page)
	return args.Error(0)
}
