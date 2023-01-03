package service

import (
	"errors"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/entity"
	pook_mock "github.com/chumnend/pook/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSrv_FindAllByBookID(t *testing.T) {
	mockRepo := new(pook_mock.MockPageRepository)
	mockPages := []entity.Page{
		{
			ID:        1,
			Content:   "page content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			BookID:    1,
		},
		{
			ID:        2,
			Content:   "page content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			BookID:    1,
		},
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindAllByBookID", mock.AnythingOfType("uint")).Return(mockPages, nil).Once()
		srv := NewPageService(mockRepo)

		// run
		pages, err := srv.FindAllByBookID(1)

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, pages, len(mockPages))
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindAllByBookID", mock.AnythingOfType("uint")).Return([]entity.Page{}, errors.New("unexpected error")).Once()
		srv := NewPageService(mockRepo)

		// run
		pages, err := srv.FindAllByBookID(1)

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, pages, 0)
		assert.Error(t, err)
	})
}

func TestSrv_FindByID(t *testing.T) {
	mockRepo := new(pook_mock.MockPageRepository)
	mockPage := entity.Page{
		ID:        1,
		Content:   "page content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		BookID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindByID", mock.AnythingOfType("uint")).Return(&mockPage, nil).Once()
		srv := NewPageService(mockRepo)

		// run
		page, err := srv.FindByID(1)

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, mockPage.ID, page.ID)
		assert.Equal(t, mockPage.Content, page.Content)
		assert.Equal(t, mockPage.BookID, page.BookID)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindByID", mock.AnythingOfType("uint")).Return(&entity.Page{}, errors.New("unexpected error")).Once()
		srv := NewPageService(mockRepo)

		// run
		page, err := srv.FindByID(1)

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, &entity.Page{}, page)
		assert.Error(t, err)
	})
}

func TestSrv_Create(t *testing.T) {
	mockRepo := new(pook_mock.MockPageRepository)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("Create", mock.Anything).Return(nil).Once()
		srv := NewPageService(mockRepo)

		// run
		err := srv.Create(&entity.Page{})

		// check
		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("Create", mock.Anything).Return(errors.New("unexpected error")).Once()
		srv := NewPageService(mockRepo)

		// run
		err := srv.Create(&entity.Page{})

		// check
		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
	})
}

func TestSrv_Update(t *testing.T) {
	mockRepo := new(pook_mock.MockPageRepository)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("Update", mock.Anything).Return(nil).Once()
		srv := NewPageService(mockRepo)

		// run
		err := srv.Update(&entity.Page{})

		// check
		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("Update", mock.Anything).Return(errors.New("unexpected error")).Once()
		srv := NewPageService(mockRepo)

		// run
		err := srv.Update(&entity.Page{})

		// check
		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
	})
}

func TestSrv_Delete(t *testing.T) {
	mockRepo := new(pook_mock.MockPageRepository)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("Delete", mock.Anything).Return(nil).Once()
		srv := NewPageService(mockRepo)

		// run
		err := srv.Delete(&entity.Page{})

		// check
		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("Delete", mock.Anything).Return(errors.New("unexpected error")).Once()
		srv := NewPageService(mockRepo)

		// run
		err := srv.Delete(&entity.Page{})

		// check
		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
	})
}

func TestSrv_Validate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		page := entity.Page{
			ID:        1,
			Content:   "test page",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			BookID:    1,
		}
		srv := NewPageService(nil)

		// run
		err := srv.Validate(&page)

		// check
		assert.Nil(t, err)
	})

	t.Run("fail - empty page", func(t *testing.T) {
		// setup
		srv := NewPageService(nil)

		// run
		err := srv.Validate(nil)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "page is empty")
	})

	t.Run("fail - missing data", func(t *testing.T) {
		// setup
		mockPage := entity.Page{
			ID:        1,
			Content:   "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			BookID:    1,
		}
		srv := NewPageService(nil)

		// run
		err := srv.Validate(&mockPage)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid page")
	})
}
