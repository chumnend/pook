package mock

import (
	"github.com/chumnend/pook/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockBookService is a mock struct of a BookService
type MockBookService struct {
	mock.Mock
}

// FindAll mock method
func (mock *MockBookService) FindAll() ([]entity.Book, error) {
	args := mock.Called()
	return args.Get(0).([]entity.Book), args.Error(1)
}

// FindAllByUserID mock method
func (mock *MockBookService) FindAllByUserID(id uint) ([]entity.Book, error) {
	args := mock.Called(id)
	return args.Get(0).([]entity.Book), args.Error(1)
}

// FindByID mock method
func (mock *MockBookService) FindByID(id uint) (*entity.Book, error) {
	args := mock.Called(id)
	return args.Get(0).(*entity.Book), args.Error(1)
}

// Create mock method
func (mock *MockBookService) Create(book *entity.Book) error {
	args := mock.Called(book)
	return args.Error(0)
}

// Save mock method
func (mock *MockBookService) Save(book *entity.Book) error {
	args := mock.Called(book)
	return args.Error(0)
}

// Delete mock method
func (mock *MockBookService) Delete(book *entity.Book) error {
	args := mock.Called(book)
	return args.Error(0)
}

// Validate mock method
func (mock *MockBookService) Validate(book *entity.Book) error {
	args := mock.Called(book)
	return args.Error(0)
}
