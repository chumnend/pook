package service

import (
	"errors"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/api/book/repository"
	"github.com/chumnend/pook/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSrv_FindAll(t *testing.T) {
	mockRepo := new(repository.MockBookRepository)
	mockBooks := []domain.Book{
		domain.Book{
			ID:        1,
			Title:     "test book 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
		domain.Book{
			ID:        2,
			Title:     "test book 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindAll").Return(mockBooks, nil).Once()
		srv := NewService(mockRepo)

		// run
		books, err := srv.FindAll()

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, books, len(mockBooks))
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindAll").Return([]domain.Book{}, errors.New("unexpected error")).Once()
		srv := NewService(mockRepo)

		// run
		books, err := srv.FindAll()

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, books, 0)
		assert.Error(t, err)
	})
}

func TestSrv_FindAllByUserID(t *testing.T) {
	mockRepo := new(repository.MockBookRepository)
	mockBooks := []domain.Book{
		domain.Book{
			ID:        1,
			Title:     "test book 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
		domain.Book{
			ID:        2,
			Title:     "test book 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindAllByUserID", mock.AnythingOfType("uint")).Return(mockBooks, nil).Once()
		srv := NewService(mockRepo)

		// run
		books, err := srv.FindAllByUserID(1)

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, books, len(mockBooks))
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindAllByUserID", mock.AnythingOfType("uint")).Return([]domain.Book{}, errors.New("unexpected error")).Once()
		srv := NewService(mockRepo)

		// run
		books, err := srv.FindAllByUserID(1)

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, books, 0)
		assert.Error(t, err)
	})
}

func TestSrv_FindByID(t *testing.T) {
	mockRepo := new(repository.MockBookRepository)
	mockBook := domain.Book{
		ID:        1,
		Title:     "test book 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindByID", mock.AnythingOfType("uint")).Return(&mockBook, nil).Once()
		srv := NewService(mockRepo)

		// run
		book, err := srv.FindByID(1)

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, mockBook.ID, book.ID)
		assert.Equal(t, mockBook.Title, book.Title)
		assert.Equal(t, mockBook.UserID, book.UserID)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindByID", mock.AnythingOfType("uint")).Return(&domain.Book{}, errors.New("unexpected error")).Once()
		srv := NewService(mockRepo)

		// run
		book, err := srv.FindByID(1)

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, &domain.Book{}, book)
		assert.Error(t, err)
	})
}

func TestSrv_Create(t *testing.T) {
	mockRepo := new(repository.MockBookRepository)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("Create", mock.Anything).Return(nil).Once()
		srv := NewService(mockRepo)

		// run
		err := srv.Create(&domain.Book{})

		// check
		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("Create", mock.Anything).Return(errors.New("unexpected error")).Once()
		srv := NewService(mockRepo)

		// run
		err := srv.Create(&domain.Book{})

		// check
		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
	})
}

func TestSrv_Save(t *testing.T) {
	mockRepo := new(repository.MockBookRepository)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("Save", mock.Anything).Return(nil).Once()
		srv := NewService(mockRepo)

		// run
		err := srv.Save(&domain.Book{})

		// check
		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("Save", mock.Anything).Return(errors.New("unexpected error")).Once()
		srv := NewService(mockRepo)

		// run
		err := srv.Save(&domain.Book{})

		// check
		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
	})
}

func TestSrv_Delete(t *testing.T) {
	mockRepo := new(repository.MockBookRepository)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("Delete", mock.Anything).Return(nil).Once()
		srv := NewService(mockRepo)

		// run
		err := srv.Delete(&domain.Book{})

		// check
		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("Delete", mock.Anything).Return(errors.New("unexpected error")).Once()
		srv := NewService(mockRepo)

		// run
		err := srv.Delete(&domain.Book{})

		// check
		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
	})
}

func TestSrv_Validate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		mockBook := domain.Book{
			ID:        1,
			Title:     "test book 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		}
		srv := NewService(nil)

		// run
		err := srv.Validate(&mockBook)

		// check
		assert.Nil(t, err)
	})

	t.Run("fail - empty book", func(t *testing.T) {
		// setup
		srv := NewService(nil)

		// run
		err := srv.Validate(nil)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "book is empty")
	})

	t.Run("fail - missing data", func(t *testing.T) {
		// setup
		mockBook := domain.Book{
			ID:        1,
			Title:     "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		}
		srv := NewService(nil)

		// run
		err := srv.Validate(&mockBook)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid book")
	})
}
