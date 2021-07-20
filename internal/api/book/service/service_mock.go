package service

import (
	"github.com/chumnend/pook/internal/api/domain"
	"github.com/stretchr/testify/mock"
)

// MockBookService is a mock struct of a BookService
type MockBookService struct {
	mock.Mock
}

// FindAll mock method
func (mock *MockBookService) FindAll() ([]domain.Book, error) {
	args := mock.Called()
	return args.Get(0).([]domain.Book), args.Error(1)
}

// FindAllByUserID mock method
func (mock *MockBookService) FindAllByUserID(id uint) ([]domain.Book, error) {
	args := mock.Called(id)
	return args.Get(0).([]domain.Book), args.Error(1)
}

// FindByID mock method
func (mock *MockBookService) FindByID(id uint) (*domain.Book, error) {
	args := mock.Called(id)
	return args.Get(0).(*domain.Book), args.Error(1)
}

// Save mock method
func (mock *MockBookService) Save(book *domain.Book) error {
	args := mock.Called(book)
	return args.Error(0)
}

// Delete mock method
func (mock *MockBookService) Delete(book *domain.Book) error {
	args := mock.Called(book)
	return args.Error(0)
}

// Validate mock method
func (mock *MockBookService) Validate(book *domain.Book) error {
	args := mock.Called(book)
	return args.Error(0)
}
