package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/book/service"
	"github.com/chumnend/pook/internal/domain"
	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func TestCtl_ListBooks(t *testing.T) {
	mockSrv := new(service.MockBookService)
	mockBooks := []domain.Book{
		domain.Book{
			ID:        1,
			Title:     "test book",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
		domain.Book{
			ID:        2,
			Title:     "test book",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
	}

	t.Run("success - find all books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAll").Return(mockBooks, nil).Once()
		ctl := NewController(mockSrv)
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books", nil)

		// run
		ctl.ListBooks(res, req)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, res.Code)
		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if _, ok := m["books"]; !ok {
			t.Errorf("Expected `books` to exist. Got '%v'", m)
			return
		}
		books := m["books"].([]interface{})
		if len(books) != 2 {
			t.Errorf("Expected 'books' to have length of 2. Got %v.", m["books"])
		}
	})

	t.Run("fail - failed to get all books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAll").Return([]domain.Book{}, errors.New("unable to access db")).Once()
		ctl := NewController(mockSrv)
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books", nil)

		// run
		ctl.ListBooks(res, req)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, res.Code)
		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["error"] != "something went wrong" {
			t.Errorf("Expected the 'error' to be 'something went wrong'. Got '%v'", m["error"])
		}
	})

	t.Run("success - find all of a user's books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByUserID", mock.AnythingOfType("uint")).Return(mockBooks, nil).Once()
		ctl := NewController(mockSrv)
		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"userID": 1}`)
		req, _ := http.NewRequest("GET", "/books", bytes.NewBuffer(jsonStr))

		// run
		ctl.ListBooks(res, req)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, res.Code)
		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if _, ok := m["books"]; !ok {
			t.Errorf("Expected `books` to exist. Got '%v'", m)
			return
		}
		books := m["books"].([]interface{})
		if len(books) != 2 {
			t.Errorf("Expected 'books' to have length of 2. Got %v.", m["books"])
		}
	})

	t.Run("fail - failed to to get user's books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByUserID", mock.AnythingOfType("uint")).Return([]domain.Book{}, errors.New("unable to access db")).Once()
		ctl := NewController(mockSrv)
		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"userID": 1}`)
		req, _ := http.NewRequest("GET", "/books", bytes.NewBuffer(jsonStr))

		// run
		ctl.ListBooks(res, req)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, res.Code)
		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if m["error"] != "something went wrong" {
			t.Errorf("Expected the 'error' to be 'something went wrong'. Got '%v'", m["error"])
		}
	})
}

func TestCtl_CreateBook(t *testing.T) {
	mockSrv := new(service.MockBookService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Save", mock.Anything).Return(nil).Once()
		ctl := NewController(mockSrv)
		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test", "userID": "1"}`)
		req, _ := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		// run
		ctl.CreateBook(res, req)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, res.Code)
		var m map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &m)
		if _, ok := m["result"]; !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		result := m["result"].(map[string]interface{})
		if result["title"] != "test" {
			t.Errorf("Expected 'title' to be 'test'. Got '%v'", m["title"])
		}
		if result["userID"] != 1.0 {
			t.Errorf("Expected `id` to be '1'. Got '%v'", m["id"])
		}
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		// check
		// run
	})
}

func TestCtl_GetBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		// run
		// check
	})
	t.Run("fail", func(t *testing.T) {
		// setup
		// check
		// run
	})
}

func TestCtl_UpdateBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		// run
		// check
	})
	t.Run("fail", func(t *testing.T) {
		// setup
		// check
		// run
	})
}

func TestCtl_DeleteBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		// run
		// check
	})
	t.Run("fail", func(t *testing.T) {
		// setup
		// check
		// run
	})
}
