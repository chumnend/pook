package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyUser(t *testing.T) {
	testService := NewService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "User is empty")
}
