package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chumnend/pook/internal/api/domain"
	"github.com/chumnend/pook/internal/api/user/mocks"
	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func TestCtl_Register(t *testing.T) {
	mockSrv := new(mocks.UserService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Save", mock.Anything).Return(nil).Once()
		mockSrv.On("GenerateToken", mock.Anything).Return("token", nil).Once()

		testController := NewController(mockSrv)

		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"email":"test@example.com", "password": "test123"}`)
		req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		// run
		testController.Register(res, req)

		// check
		mockSrv.AssertExpectations(t)

		checkResponseCode(t, http.StatusOK, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["token"] == "" {
			t.Errorf("Expected 'token' to be non empty. Got %v.", m["token"])
		}
	})

	t.Run("fail - no email", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(errors.New("missing and/or invalid information")).Once()

		testController := NewController(mockSrv)

		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"password": "test123"}`)
		req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		// run
		testController.Register(res, req)

		// check
		mockSrv.AssertExpectations(t)
		mockSrv.AssertNotCalled(t, "Save")
		mockSrv.AssertNotCalled(t, "GenerateToken")

		checkResponseCode(t, http.StatusBadRequest, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)

		if m["error"] != "missing and/or invalid information" {
			t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - no password", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(errors.New("missing and/or invalid information")).Once()

		testController := NewController(mockSrv)

		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"email":"test@example.com"}`)
		req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		// run
		testController.Register(res, req)

		// check
		mockSrv.AssertExpectations(t)
		mockSrv.AssertNotCalled(t, "Save")
		mockSrv.AssertNotCalled(t, "GenerateToken")

		checkResponseCode(t, http.StatusBadRequest, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["error"] != "missing and/or invalid information" {
			t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
		}
	})
}

func TestCtl_Login(t *testing.T) {
	mockSrv := new(mocks.UserService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByEmail", mock.AnythingOfType("string")).Return(&domain.User{}, nil).Once()
		mockSrv.On("ComparePassword", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
		mockSrv.On("GenerateToken", mock.Anything).Return("token", nil).Once()

		testController := NewController(mockSrv)

		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
		req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		// run
		testController.Login(res, req)

		// check
		mockSrv.AssertExpectations(t)

		checkResponseCode(t, http.StatusOK, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["token"] == "" {
			t.Errorf("Expected 'token' to be non empty. Got %v.", m["token"])
		}
	})

	t.Run("fail - bad email", func(t *testing.T) {
		// setup
		mockSrv.On("FindByEmail", mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("invalid email and/or password")).Once()

		testController := NewController(mockSrv)

		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
		req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		// run
		testController.Login(res, req)

		// check
		mockSrv.AssertExpectations(t)
		mockSrv.AssertNotCalled(t, "ComparePassword")
		mockSrv.AssertNotCalled(t, "GenerateToken")

		checkResponseCode(t, http.StatusBadRequest, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["error"] != "invalid email and/or password" {
			t.Errorf("Expected the 'error' to be 'invalid email and/or password'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - bad password", func(t *testing.T) {
		// setup
		mockSrv.On("FindByEmail", mock.AnythingOfType("string")).Return(&domain.User{}, nil).Once()
		mockSrv.On("ComparePassword", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("invalid email and/or password")).Once()

		testController := NewController(mockSrv)

		res := httptest.NewRecorder()
		jsonStr := []byte(`{"email":"test@example.com", "password": "123"}`)
		req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		// run
		testController.Login(res, req)

		// check
		mockSrv.AssertExpectations(t)
		mockSrv.AssertNotCalled(t, "GenerateToken")

		checkResponseCode(t, http.StatusBadRequest, res.Code)

		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["error"] != "invalid email and/or password" {
			t.Errorf("Expected the 'error' to be 'invalid email and/or password'. Got '%v'", m["error"])
		}
	})
}
