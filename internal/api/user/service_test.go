package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/chumnend/pook/internal/api/domain"
	"github.com/chumnend/pook/internal/api/user/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSrv_FindAll(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		Email:    "tester@pook.com",
		Password: "123",
	}

	mockListUsers := make([]domain.User, 0)
	mockListUsers = append(mockListUsers, mockUser)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindAll").Return(mockListUsers, nil).Once()
		testService := NewService(mockRepo)

		// run
		result, err := testService.FindAll()

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, result, len(mockListUsers))
		assert.Equal(t, mockUser.Email, result[0].Email)
		assert.Equal(t, mockUser.Password, result[0].Password)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindAll").Return([]domain.User{}, errors.New("unexpected error")).Once()
		testService := NewService(mockRepo)

		// run
		result, err := testService.FindAll()

		// check
		mockRepo.AssertExpectations(t)
		assert.Len(t, result, 0)
		assert.Error(t, err)
	})
}

func TestSrv_FindByEmail(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		Email:    "tester@pook.com",
		Password: "123",
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(&mockUser, nil).Once()
		testService := NewService(mockRepo)

		// run
		result, err := testService.FindByEmail("tester@pook.com")

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, mockUser.Email, result.Email)
		assert.Equal(t, mockUser.Password, result.Password)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("unexpected error")).Once()
		testService := NewService(mockRepo)

		// run
		result, err := testService.FindByEmail("tester@pook.com")

		// check
		mockRepo.AssertExpectations(t)
		assert.Equal(t, "", result.Email)
		assert.Equal(t, "", result.Password)
		assert.Error(t, err)
	})
}

func TestSrv_Save(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {
		// setup
		mockRepo.On("Save", mock.Anything).Return(nil).Once()
		testService := NewService(mockRepo)

		// run
		result := testService.Save(&domain.User{})

		// check
		mockRepo.AssertExpectations(t)
		assert.NoError(t, result)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		mockRepo.On("Save", mock.Anything).Return(errors.New("unexpected error")).Once()
		testService := NewService(mockRepo)

		// run
		result := testService.Save(&domain.User{})

		// check
		mockRepo.AssertExpectations(t)
		assert.Error(t, result)
	})
}

func TestSrv_Validate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		testService := NewService(nil)
		mockUser := &domain.User{
			Email:    "tester@pook.com",
			Password: "123",
		}

		// run
		err := testService.Validate(mockUser)

		// check
		assert.Nil(t, err)
	})

	t.Run("fail - empty user", func(t *testing.T) {
		// setup
		testService := NewService(nil)

		// run
		err := testService.Validate(nil)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "user is empty")
	})

	t.Run("fail - no email", func(t *testing.T) {
		// setup
		testService := NewService(nil)
		mockUser := &domain.User{
			Email:    "",
			Password: "123",
		}

		// run
		err := testService.Validate(mockUser)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid user")
	})

	t.Run("fail - no password", func(t *testing.T) {
		// setup
		testService := NewService(nil)
		mockUser := &domain.User{
			Email:    "tester@pook.com",
			Password: "",
		}

		// run
		err := testService.Validate(mockUser)

		// check
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid user")
	})
}

func TestSrv_GenerateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		testService := NewService(nil)
		mockUser := &domain.User{
			Email:    "tester@pook.com",
			Password: "123",
		}

		// run
		token, err := testService.GenerateToken(mockUser)

		// check
		assert.Greater(t, len(token), 0)
		assert.NoError(t, err)
	})
}

func TestSrv_ComparePassword(t *testing.T) {
	testService := NewService(nil)
	password := "123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	mockUser := &domain.User{
		Email:    "tester@pook.com",
		Password: string(hashedPassword),
	}

	t.Run("success", func(t *testing.T) {
		// run
		err := testService.ComparePassword(mockUser, password)

		// check
		assert.NoError(t, err)
	})

	t.Run("fail - wrong password", func(t *testing.T) {
		// run
		err := testService.ComparePassword(mockUser, "")

		// check
		assert.Error(t, err)
	})
}
