package mock

import (
	"github.com/chumnend/pook/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockPageRepository is a mock for the BookRepository
type MockPageRepository struct {
	mock.Mock
}

// FindAllByBookID mock method
func (mock *MockPageRepository) FindAllByBookID(id uint) ([]entity.Page, error) {
	args := mock.Called(id)
	return args.Get(0).([]entity.Page), args.Error(1)
}

// FindByID mock method
func (mock *MockPageRepository) FindByID(id uint) (*entity.Page, error) {
	args := mock.Called(id)
	return args.Get(0).(*entity.Page), args.Error(1)
}

// Create mock method
func (mock *MockPageRepository) Create(page *entity.Page) error {
	args := mock.Called(page)
	return args.Error(0)
}

// Update mock method
func (mock *MockPageRepository) Update(page *entity.Page) error {
	args := mock.Called(page)
	return args.Error(0)
}

// Delete mock method
func (mock *MockPageRepository) Delete(page *entity.Page) error {
	args := mock.Called(page)
	return args.Error(0)
}

// Migrate mock method
func (mock *MockPageRepository) Migrate() error {
	args := mock.Called()
	return args.Error(0)
}
