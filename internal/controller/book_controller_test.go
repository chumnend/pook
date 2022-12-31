package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/entity"
	pook_mock "github.com/chumnend/pook/internal/mock"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestBookController_ListBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSrv := new(pook_mock.MockBookService)
	mockBooks := []entity.Book{
		{
			ID:        1,
			Title:     "test book",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
		{
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
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/books", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.ListBooks(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["books"]
		if !ok {
			t.Errorf("Expected `books` to exist. Got '%v'", m)
			return
		}
		books := value.([]interface{})
		if len(books) != 2 {
			t.Errorf("Expected 'books' to have length of 2. Got %v.", len(books))
		}
	})

	t.Run("fail - failed to get all books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAll").Return([]entity.Book{}, errors.New("unable to access db")).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/books", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.ListBooks(c)

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

	t.Run("success - find all of a user's books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByUserID", mock.AnythingOfType("uint")).Return(mockBooks, nil).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/books?userId=1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.ListBooks(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["books"]
		if !ok {
			t.Errorf("Expected `books` to exist. Got '%v'", m)
			return
		}
		books := value.([]interface{})
		if len(books) != 2 {
			t.Errorf("Expected 'books' to have length of 2. Got %v.", len(books))
		}
	})

	t.Run("fail - failed to to get user's books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByUserID", mock.AnythingOfType("uint")).Return([]entity.Book{}, errors.New("unable to access db")).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/books?userId=1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.ListBooks(c)

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

func TestBookController_CreateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSrv := new(pook_mock.MockBookService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Create", mock.Anything).Return(nil).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test", "userID": "1"}`)
		r, _ := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.CreateBook(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["book"]
		if !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		book := value.(map[string]interface{})
		if book["title"] != "test" {
			t.Errorf("Expected 'title' to be 'test'. Got '%v'", book["title"])
		}
		if book["userID"] != 1.0 {
			t.Errorf("Expected `userID` to be '1'. Got '%v'", book["userID"])
		}
	})

	t.Run("fail - missing userID in request", func(t *testing.T) {
		// setup
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test"}`)
		r, _ := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.CreateBook(c)

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

	t.Run("fail - unable to validate book struct", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test", "userID": "123"}`)
		r, _ := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.CreateBook(c)

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

	t.Run("fail - unable to save", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Create", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test", "userID": "1"}`)
		r, _ := http.NewRequest("POST", "/v1/books", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.CreateBook(c)

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

func TestBookController_GetBook(t *testing.T) {
	mockSrv := new(pook_mock.MockBookService)
	mockBook := entity.Book{
		ID:        1,
		Title:     "test book",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockBook, nil).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/books/1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.GetBook(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["book"]
		if !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		book := value.(map[string]interface{})
		assert.Equal(t, float64(mockBook.ID), book["id"]) // FixMe: Hacky comparison of uint
		assert.Equal(t, mockBook.Title, book["title"])
		assert.Equal(t, float64(mockBook.UserID), book["userID"]) // FixMe: Hacky comparison of uint
	})

	t.Run("fail - invalid book id", func(t *testing.T) {
		// setup
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/books/abc", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "abc",
			},
		}

		// run
		ctl.GetBook(c)

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

	t.Run("fail - book not found", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&entity.Book{}, errors.New("unexpected error")).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/books/1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.GetBook(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusNotFound, w.Code)
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

func TestBookController_UpdateBook(t *testing.T) {
	mockSrv := new(pook_mock.MockBookService)
	mockBook := entity.Book{
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
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/books/1", bytes.NewBuffer(jsonStr))
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.UpdateBook(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["book"]
		if !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		book := value.(map[string]interface{})
		if book["title"] != "test" {
			t.Errorf("Expected 'title' to be 'test'. Got '%v'", book["title"])
		}
		if book["userID"] != 1.0 {
			t.Errorf("Expected `userID` to be '1'. Got '%v'", book["userID"])
		}
	})

	t.Run("fail - save error", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockBook, nil).Once()
		mockSrv.On("Save", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/books/1", bytes.NewBuffer(jsonStr))
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.UpdateBook(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
	})
}

func TestBookController_DeleteBook(t *testing.T) {
	mockSrv := new(pook_mock.MockBookService)
	mockBook := entity.Book{
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
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/books/1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.DeleteBook(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
	})

	t.Run("fail - delete error", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockBook, nil).Once()
		mockSrv.On("Delete", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewBookController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/books/1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.DeleteBook(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
	})
}
