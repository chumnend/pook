package user

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/chumnend/pook/internal/pook/domain"
	"github.com/chumnend/pook/internal/pook/user/mocks"
	"github.com/stretchr/testify/assert"
)

// FindAll
func TestFindAll(t *testing.T) {
	// setup
	mockRepo := new(mocks.UserRepository)
	user := domain.User{Email: "tester@pook.com", Password: "123"}
	mockRepo.On("FindAll").Return([]domain.User{user}, nil)
	testService := NewService(mockRepo)

	// run
	result, _ := testService.FindAll()

	// check
	mockRepo.AssertExpectations(t)
	assert.Equal(t, user.Email, result[0].Email)
	assert.Equal(t, user.Password, result[0].Password)
}

// FindByEmail
func TestFindByEmail(t *testing.T) {
	// setup
	mockRepo := new(mocks.UserRepository)
	user := domain.User{Email: "tester@pook.com", Password: "123"}
	mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(&user, nil)
	testService := NewService(mockRepo)

	// run
	result, _ := testService.FindByEmail("tester@pook.com")

	// check
	mockRepo.AssertExpectations(t)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Password, result.Password)
}

// Validate
func TestValidateUser(t *testing.T) {
	// setup
	testService := NewService(nil)
	user := &domain.User{Email: "tester@pook.com", Password: "123"}

	// run
	err := testService.Validate(user)

	// check
	assert.Nil(t, err)
}

func TestValidateEmptyUser(t *testing.T) {
	// setup
	testService := NewService(nil)

	// run
	err := testService.Validate(nil)

	// check
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "User is empty")
}

func TestValidateEmptyEmail(t *testing.T) {
	// setup
	testService := NewService(nil)
	user := &domain.User{Email: "", Password: "123"}

	// run
	err := testService.Validate(user)

	// check
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Invalid User")
}

func TestValidateEmptyPassword(t *testing.T) {
	// setup
	testService := NewService(nil)
	user := &domain.User{Email: "tester@pook.com", Password: ""}

	// run
	err := testService.Validate(user)

	// check
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Invalid User")
}
