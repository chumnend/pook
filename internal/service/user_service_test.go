package service

import (
	"errors"
	"testing"

	"github.com/chumnend/pook/internal/entity"
	pook_mock "github.com/chumnend/pook/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_FindAll(t *testing.T) {
	mockRepo := new(pook_mock.MockUserRepository)
	mockUsers := []entity.User{
		{
			Email:     "tester@pook.com",
			FirstName: "tester",
			LastName:  "tester",
			Password:  "123",
		},
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindAll").Return(mockUsers, nil).Once()
		srv := NewUserService(mockRepo)

		// run
		result, err := srv.FindAll()

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, result, len(mockUsers))
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindAll").Return([]entity.User{}, errors.New("unexpected error")).Once()
		srv := NewUserService(mockRepo)

		// run
		users, err := srv.FindAll()

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, users, 0)
		assert.Error(t, err)
	})
}

func TestUserService_FindByUsername(t *testing.T) {
	mockRepo := new(pook_mock.MockUserRepository)
	mockUser := entity.User{
		Username:  "tester",
		Email:     "tester@pook.com",
		FirstName: "tester",
		LastName:  "tester",
		Password:  "123",
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindByUsername", mock.AnythingOfType("string")).Return(&mockUser, nil).Once()
		srv := NewUserService(mockRepo)

		// run
		user, err := srv.FindByUsername("tester")

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, mockUser.Username, user.Username)
		assert.Equal(t, mockUser.Email, user.Email)
		assert.Equal(t, mockUser.Password, user.Password)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindByUsername", mock.AnythingOfType("string")).Return(&entity.User{}, errors.New("unexpected error")).Once()
		srv := NewUserService(mockRepo)

		// run
		user, err := srv.FindByUsername("tester")

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, &entity.User{}, user)
		assert.Error(t, err)
	})
}

func TestUserService_FindByEmail(t *testing.T) {
	mockRepo := new(pook_mock.MockUserRepository)
	mockUser := entity.User{
		Username:  "tester",
		Email:     "tester@pook.com",
		FirstName: "tester",
		LastName:  "tester",
		Password:  "123",
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(&mockUser, nil).Once()
		srv := NewUserService(mockRepo)

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
		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(&entity.User{}, errors.New("unexpected error")).Once()
		srv := NewUserService(mockRepo)

		// run
		user, err := srv.FindByEmail("tester@pook.com")

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, &entity.User{}, user)
		assert.Error(t, err)
	})
}

func TestUserService_Save(t *testing.T) {
	mockRepo := new(pook_mock.MockUserRepository)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("Save", mock.Anything).Return(nil).Once()
		srv := NewUserService(mockRepo)

		// run
		err := srv.Save(&entity.User{})

		// check
		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("Save", mock.Anything).Return(errors.New("unexpected error")).Once()
		srv := NewUserService(mockRepo)

		// run
		err := srv.Save(&entity.User{})

		// check
		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
	})
}

func TestUserService_GenerateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		srv := NewUserService(nil)
		mockUser := entity.User{
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

func TestUserService_ComparePassword(t *testing.T) {
	srv := NewUserService(nil)
	password := "123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	mockUser := &entity.User{
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
