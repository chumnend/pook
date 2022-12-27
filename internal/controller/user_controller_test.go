package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chumnend/pook/internal/entity"
	pook_mock "github.com/chumnend/pook/internal/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func TestCtl_Register(t *testing.T) {
	mockSrv := new(pook_mock.MockUserService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("Save", mock.Anything).Return(nil).Once()
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"username": "test", "email":"test@example.com", "password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Register(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["message"]
		if !ok {
			t.Errorf("Unable to find key 'message'.")
		}
		if value == "" {
			t.Errorf("Expected 'message' to be non empty. Got %v.", value)
		}
	})

	t.Run("fail - no username", func(t *testing.T) {
		// setup
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"email": "test@example.com", ""password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Register(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["error"]
		if !ok {
			t.Errorf("Unable to find key 'error'.")
		}
		if value == "" {
			t.Errorf("Expected 'error' to be non empty. Got %v.", value)
		}
	})

	t.Run("fail - no email", func(t *testing.T) {
		// setup
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"username": "test", "password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Register(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["error"]
		if !ok {
			t.Errorf("Unable to find key 'error'.")
		}
		if value == "" {
			t.Errorf("Expected 'error' to be non empty. Got %v.", value)
		}
	})

	t.Run("fail - no password", func(t *testing.T) {
		// setup
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"username": "test", "email":"test@example.com"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Register(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["error"]
		if !ok {
			t.Errorf("Unable to find key 'error'.")
		}
		if value == "" {
			t.Errorf("Expected 'error' to be non empty. Got %v.", value)
		}
	})

	t.Run("fail - save error", func(t *testing.T) {
		// setup
		mockSrv.On("Save", mock.Anything).Return(errors.New("save error")).Once()
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"username": "test", "email":"test@example.com", "password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Register(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["error"]
		if !ok {
			t.Errorf("Unable to find key 'error'.")
		}
		if value == "" {
			t.Errorf("Expected 'error' to be non empty. Got %v.", value)
		}
	})
}

func TestCtl_Login(t *testing.T) {
	mockSrv := new(pook_mock.MockUserService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByUsername", mock.AnythingOfType("string")).Return(&entity.User{}, nil).Once()
		mockSrv.On("ComparePassword", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
		mockSrv.On("GenerateToken", mock.Anything).Return("token", nil).Once()
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"username": "test", "password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Login(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["token"]
		if !ok {
			t.Errorf("Unable to find key 'token'.")
		}
		if value == "" {
			t.Errorf("Expected 'token' to be non empty. Got %v.", value)
		}
	})

	t.Run("fail - bad input", func(t *testing.T) {
		// setup
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"username": "test"}`)
		r, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Login(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["error"]
		if !ok {
			t.Errorf("Unable to find key 'error'.")
		}
		if value == "" {
			t.Errorf("Expected 'error' to be non empty. Got %v.", value)
		}
	})

	t.Run("fail - bad username", func(t *testing.T) {
		// setup
		mockSrv.On("FindByUsername", mock.AnythingOfType("string")).Return(&entity.User{}, errors.New("invalid email and/or password")).Once()
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"username": "test", "password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Login(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["error"]
		if !ok {
			t.Errorf("Unable to find key 'error'.")
		}
		if value == "" {
			t.Errorf("Expected 'error' to be non empty. Got %v.", value)
		}
	})

	t.Run("fail - bad password", func(t *testing.T) {
		// setup
		mockSrv.On("FindByUsername", mock.AnythingOfType("string")).Return(&entity.User{}, nil).Once()
		mockSrv.On("ComparePassword", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("invalid email and/or password")).Once()
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"username": "test", "password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Login(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["error"]
		if !ok {
			t.Errorf("Unable to find key 'error'.")
		}
		if value == "" {
			t.Errorf("Expected 'error' to be non empty. Got %v.", value)
		}
	})

	t.Run("fail - token error", func(t *testing.T) {
		// setup
		mockSrv.On("FindByUsername", mock.AnythingOfType("string")).Return(&entity.User{}, nil).Once()
		mockSrv.On("ComparePassword", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
		mockSrv.On("GenerateToken", mock.Anything).Return("", errors.New("token error")).Once()
		ctl := NewController(mockSrv)
		var jsonStr = []byte(`{"username": "test", "password": "test123"}`)
		r, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.Login(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["error"]
		if !ok {
			t.Errorf("Unable to find key 'error'.")
		}
		if value == "" {
			t.Errorf("Expected 'error' to be non empty. Got %v.", value)
		}
	})
}
