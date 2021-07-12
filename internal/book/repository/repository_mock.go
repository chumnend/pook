package repository

import (
	"github.com/chumnend/pook/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockBookRepository is a mock for the BookRepository
type MockBookRepository struct {
	mock.Mock
}

// FindAll method mock
func (mock *MockBookRepository) FindAll() ([]domain.Book, error) {
	args := mock.Called()
	return args.Get(0).([]domain.Book), args.Error(1)
}

// FindAllByUserID method mock
func (mock *MockBookRepository) FindAllByUserID(id uint) ([]domain.Book, error) {
	args := mock.Called(id)
	return args.Get(0).([]domain.Book), args.Error(1)
}

// FindByID method mock
func (mock *MockBookRepository) FindByID(id uint) (*domain.Book, error) {
	args := mock.Called(id)
	return args.Get(0).(*domain.Book), args.Error(1)
}

// Save method mock
func (mock *MockBookRepository) Save(book *domain.Book) error {
	args := mock.Called(book)
	return args.Error(0)
}

// Delete method mock
func (mock *MockBookRepository) Delete(book *domain.Book) error {
	args := mock.Called(book)
	return args.Error(0)
}

// Migrate method mock
func (mock *MockBookRepository) Migrate() error {
	args := mock.Called()
	return args.Error(0)
}
