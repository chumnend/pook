package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/chumnend/pook/internal/domain"
	"github.com/chumnend/pook/internal/user/repository"
	"github.com/stretchr/testify/assert"
)

func TestSrv_FindAll(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	mockUsers := []domain.User{
		domain.User{
			Email:     "tester@pook.com",
			FirstName: "tester",
			LastName:  "tester",
			Password:  "123",
		},
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindAll").Return(mockUsers, nil).Once()
		srv := NewService(mockRepo)

		// run
		result, err := srv.FindAll()

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, result, len(mockUsers))
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindAll").Return([]domain.User{}, errors.New("unexpected error")).Once()
		srv := NewService(mockRepo)

		// run
		users, err := srv.FindAll()

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, users, 0)
		assert.Error(t, err)
	})
}

func TestSrv_FindByEmail(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	mockUser := domain.User{
		Email:     "tester@pook.com",
		FirstName: "tester",
		LastName:  "tester",
		Password:  "123",
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(&mockUser, nil).Once()
		srv := NewService(mockRepo)

		// run
		user, err := srv.FindByEmail("tester@pook.com")

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, mockUser.Email, user.Email)
		assert.Equal(t, mockUser.Password, user.Password)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("unexpected error")).Once()
		srv := NewService(mockRepo)

		// run
		user, err := srv.FindByEmail("tester@pook.com")

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, &domain.User{}, user)
		assert.Error(t, err)
	})
}

func TestSrv_Save(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("Save", mock.Anything).Return(nil).Once()
		srv := NewService(mockRepo)

		// run
		err := srv.Save(&domain.User{})

		// check
		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("Save", mock.Anything).Return(errors.New("unexpected error")).Once()
		srv := NewService(mockRepo)

		// run
		err := srv.Save(&domain.User{})

		// check
		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
	})
}

func TestSrv_Validate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		srv := NewService(nil)
		mockUser := domain.User{
			Email:     "tester@pook.com",
			FirstName: "tester",
			LastName:  "tester",
			Password:  "123",
		}

		// run
		err := srv.Validate(&mockUser)

		// check
		assert.Nil(t, err)
	})

	t.Run("fail - empty user", func(t *testing.T) {
		// setup
		srv := NewService(nil)

		// run
		err := srv.Validate(nil)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "user is empty")
	})

	t.Run("fail - no email", func(t *testing.T) {
		// setup
		srv := NewService(nil)
		mockUser := domain.User{
			Email:     "",
			FirstName: "tester",
			LastName:  "tester",
			Password:  "123",
		}

		// run
		err := srv.Validate(&mockUser)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid user")
	})

	t.Run("fail - no password", func(t *testing.T) {
		// setup
		srv := NewService(nil)
		mockUser := domain.User{
			Email:     "tester@pook.com",
			FirstName: "tester",
			LastName:  "tester",
			Password:  "",
		}

		// run
		err := srv.Validate(&mockUser)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid user")
	})
}

func TestSrv_GenerateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		srv := NewService(nil)
		mockUser := domain.User{
			Email:     "tester@pook.com",
			FirstName: "tester",
			LastName:  "tester",
			Password:  "123",
		}

		// run
		token, err := srv.GenerateToken(&mockUser)

		// check
		assert.Greater(t, len(token), 0)
		assert.NoError(t, err)
	})
}

func TestSrv_ComparePassword(t *testing.T) {
	srv := NewService(nil)
	password := "123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	mockUser := &domain.User{
		Email:     "tester@pook.com",
		FirstName: "tester",
		LastName:  "tester",
		Password:  string(hashedPassword),
	}

	t.Run("success", func(t *testing.T) {
		// run
		err := srv.ComparePassword(mockUser, password)

		// check
		assert.NoError(t, err)
	})

	t.Run("fail - wrong password", func(t *testing.T) {
		// run
		err := srv.ComparePassword(mockUser, "")

		// check
		assert.Error(t, err)
	})
}
