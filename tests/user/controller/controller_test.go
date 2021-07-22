package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chumnend/pook/internal/domain"
	"github.com/chumnend/pook/internal/user/controller"
	"github.com/chumnend/pook/tests/user/service"
	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func TestCtl_Register(t *testing.T) {
	mockSrv := new(service.MockUserService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Save", mock.Anything).Return(nil).Once()
		mockSrv.On("GenerateToken", mock.Anything).Return("token", nil).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"email":"test@example.com", "password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.Register(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["token"] == "" {
			t.Errorf("Expected 'token' to be non empty. Got %v.", m["token"])
		}
	})

	t.Run("fail - no email", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(errors.New("missing and/or invalid information")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.Register(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "missing and/or invalid information" {
			t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - no password", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(errors.New("missing and/or invalid information")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"email":"test@example.com"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.Register(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "missing and/or invalid information" {
			t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
		}
	})
}

func TestCtl_Login(t *testing.T) {
	mockSrv := new(service.MockUserService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByEmail", mock.AnythingOfType("string")).Return(&domain.User{}, nil).Once()
		mockSrv.On("ComparePassword", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
		mockSrv.On("GenerateToken", mock.Anything).Return("token", nil).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.Login(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["token"] == "" {
			t.Errorf("Expected 'token' to be non empty. Got %v.", m["token"])
		}
	})

	t.Run("fail - bad email", func(t *testing.T) {
		// setup
		mockSrv.On("FindByEmail", mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("invalid email and/or password")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"email":"test@example.com", "password": "123"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.Login(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "invalid email and/or password" {
			t.Errorf("Expected the 'error' to be 'invalid email and/or password'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - bad password", func(t *testing.T) {
		// setup
		mockSrv.On("FindByEmail", mock.AnythingOfType("string")).Return(&domain.User{}, nil).Once()
		mockSrv.On("ComparePassword", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("invalid email and/or password")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		jsonStr := []byte(`{"email":"test@example.com", "password": "123"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.Login(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "invalid email and/or password" {
			t.Errorf("Expected the 'error' to be 'invalid email and/or password'. Got '%v'", m["error"])
		}
	})
}
