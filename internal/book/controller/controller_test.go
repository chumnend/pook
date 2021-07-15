package controller

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/book/service"
	"github.com/chumnend/pook/internal/domain"
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
		mockSrv.AssertNotCalled(t, "FindAllByUserID")
		checkResponseCode(t, http.StatusOK, res.Code)
	})

	t.Run("success - find all of a particular user's books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByUserID").Return(mockBooks, nil).Once()
		ctl := NewController(mockSrv)
		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"user_id": 1}`)
		req, _ := http.NewRequest("GET", "/books", bytes.NewBuffer(jsonStr))

		// run
		ctl.ListBooks(res, req)

		// check
		mockSrv.AssertExpectations(t)
		mockSrv.AssertNotCalled(t, "FindAll")
		checkResponseCode(t, http.StatusOK, res.Code)
	})

	t.Run("fail - failed to get books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAll").Return([]domain.Book{}, errors.New("unable to access db")).Once()
		ctl := NewController(mockSrv)
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books", nil)

		// run
		ctl.ListBooks(res, req)

		// check
		mockSrv.AssertExpectations(t)
		mockSrv.AssertNotCalled(t, "FindAllByUserID")
		checkResponseCode(t, http.StatusBadRequest, res.Code)
	})

	t.Run("fail - failed to to get books", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByUserID").Return([]domain.Book{}, errors.New("unable to access db")).Once()
		ctl := NewController(mockSrv)
		res := httptest.NewRecorder()
		var jsonStr = []byte(`{"user_id": 1}`)
		req, _ := http.NewRequest("GET", "/books", bytes.NewBuffer(jsonStr))

		// run
		ctl.ListBooks(res, req)

		// check
		mockSrv.AssertExpectations(t)
		mockSrv.AssertNotCalled(t, "FindAll")
		checkResponseCode(t, http.StatusBadRequest, res.Code)
	})
}

func TestCtl_CreateBook(t *testing.T) {
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
