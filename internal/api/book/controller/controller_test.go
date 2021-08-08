package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/api/book/service"
	"github.com/chumnend/pook/internal/domain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
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
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/books", nil)

		// run
		ctl.ListBooks(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
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
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/books", nil)

		// run
		ctl.ListBooks(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "something went wrong" {
			t.Errorf("Expected the 'error' to be 'something went wrong'. Got '%v'", m["error"])
		}
	})

	t.Run("success - find all of a user's books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByUserID", mock.AnythingOfType("uint")).Return(mockBooks, nil).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/books?userId=1", nil)

		// run
		ctl.ListBooks(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
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
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/books?userId=1", nil)

		// run
		ctl.ListBooks(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
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
		mockSrv.On("Create", mock.Anything).Return(nil).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test", "userID": "1"}`)
		r, _ := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.CreateBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if _, ok := m["book"]; !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		result := m["book"].(map[string]interface{})
		if result["title"] != "test" {
			t.Errorf("Expected 'title' to be 'test'. Got '%v'", m["title"])
		}
		if result["userID"] != 1.0 {
			t.Errorf("Expected `userID` to be '1'. Got '%v'", m["userID"])
		}
	})

	t.Run("fail - bad book", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test"}`)
		r, _ := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.CreateBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "missing and/or invalid information" {
			t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - bad save", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Create", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test", "userID": "1"}`)
		r, _ := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.CreateBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "something went wrong" {
			t.Errorf("Expected the 'error' to be 'something went wrong'. Got '%v'", m["error"])
		}
	})
}

func TestCtl_GetBook(t *testing.T) {
	mockSrv := new(service.MockBookService)
	mockBook := domain.Book{
		ID:        1,
		Title:     "test book",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockBook, nil).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/books/1", nil)
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.GetBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if _, ok := m["book"]; !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		result := m["book"].(map[string]interface{})
		assert.Equal(t, float64(mockBook.ID), result["id"]) // FixMe: Hacky comparison of uint
		assert.Equal(t, mockBook.Title, result["title"])
		assert.Equal(t, float64(mockBook.UserID), result["userID"]) // FixMe: Hacky comparison of uint
	})

	t.Run("fail - invalid book id", func(t *testing.T) {
		// setup
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/books/abc", nil)
		vars := map[string]string{"id": "abc"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.GetBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "invalid book id" {
			t.Errorf("Expected the 'error' to be 'invalid book id'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - book not found", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&domain.Book{}, errors.New("unexpected error")).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/books/1", nil)
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.GetBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusNotFound, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "book not found" {
			t.Errorf("Expected the 'error' to be 'book not found'. Got '%v'", m["error"])
		}
	})
}

func TestCtl_UpdateBook(t *testing.T) {
	mockSrv := new(service.MockBookService)
	mockBook := domain.Book{
		ID:        1,
		Title:     "test book",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockBook, nil).Once()
		mockSrv.On("Save", mock.Anything).Return(nil).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/books/1", bytes.NewBuffer(jsonStr))
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.UpdateBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if _, ok := m["book"]; !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		result := m["book"].(map[string]interface{})
		if result["title"] != "test" {
			t.Errorf("Expected 'title' to be 'test'. Got '%v'", m["title"])
		}
		if result["userID"] != 1.0 {
			t.Errorf("Expected `id` to be '1'. Got '%v'", m["id"])
		}
	})

	t.Run("fail - save error", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockBook, nil).Once()
		mockSrv.On("Save", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/books/1", bytes.NewBuffer(jsonStr))
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.UpdateBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
	})
}

func TestCtl_DeleteBook(t *testing.T) {
	mockSrv := new(service.MockBookService)
	mockBook := domain.Book{
		ID:        1,
		Title:     "test book",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockBook, nil).Once()
		mockSrv.On("Delete", mock.Anything).Return(nil).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/books/1", nil)
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.DeleteBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
	})

	t.Run("fail - delete error", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockBook, nil).Once()
		mockSrv.On("Delete", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/books/1", nil)
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.DeleteBook(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
	})
}
