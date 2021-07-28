package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/domain"
	"github.com/chumnend/pook/internal/page/controller"
	"github.com/chumnend/pook/internal/page/service"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func TestCtl_ListPages(t *testing.T) {
	mockSrv := new(service.MockPageService)
	mockPages := []domain.Page{
		domain.Page{
			ID:        1,
			Content:   "page content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			BookID:    1,
		},
		domain.Page{
			ID:        2,
			Content:   "page content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			BookID:    1,
		},
	}

	t.Run("success - find all of a book's pages", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByBookID", mock.AnythingOfType("uint")).Return(mockPages, nil).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/pages?bookId=1", nil)

		// run
		ctl.ListPages(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if _, ok := m["pages"]; !ok {
			t.Errorf("Expected `pages` to exist. Got '%v'", m)
			return
		}
		pages := m["pages"].([]interface{})
		if len(pages) != 2 {
			t.Errorf("Expected 'pages' to have length of 2. Got %v.", m["pages"])
		}
	})

	t.Run("fail - failed to to get book's pages", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByBookID", mock.AnythingOfType("uint")).Return([]domain.Page{}, errors.New("unable to access db")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/pages?bookId=1", nil)

		// run
		ctl.ListPages(w, r)

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

func TestCtl_CreatePage(t *testing.T) {
	mockSrv := new(service.MockPageService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Create", mock.Anything).Return(nil).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test", "bookID": "1"}`)
		r, _ := http.NewRequest("POST", "/v1/pages", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.CreatePage(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if _, ok := m["result"]; !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		result := m["result"].(map[string]interface{})
		if result["content"] != "test" {
			t.Errorf("Expected 'content' to be 'test'. Got '%v'", m["content"])
		}
		if result["bookID"] != 1.0 {
			t.Errorf("Expected `bookID` to be '1'. Got '%v'", m["bookID"])
		}
	})

	t.Run("fail - bad page", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test"}`)
		r, _ := http.NewRequest("POST", "/v1/pages", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.CreatePage(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "missing and/or invalid information" {
			t.Errorf("Expected the 'error' to be 'missing and/or invalid information'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - bad create", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Create", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test", "bookID": "1"}`)
		r, _ := http.NewRequest("POST", "/v1/pages", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// run
		ctl.CreatePage(w, r)

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

func TestCtl_GetPage(t *testing.T) {
	mockSrv := new(service.MockPageService)
	mockPage := domain.Page{
		ID:        1,
		Content:   "page content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		BookID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockPage, nil).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/pages/1", nil)
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.GetPage(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if _, ok := m["result"]; !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		result := m["result"].(map[string]interface{})
		assert.Equal(t, float64(mockPage.ID), result["id"]) // FixMe: Hacky comparison of uint
		assert.Equal(t, mockPage.Content, result["content"])
		assert.Equal(t, float64(mockPage.BookID), result["bookID"]) // FixMe: Hacky comparison of uint
	})

	t.Run("fail - invalid page id", func(t *testing.T) {
		// setup
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/pages/abc", nil)
		vars := map[string]string{"id": "abc"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.GetPage(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "invalid page id" {
			t.Errorf("Expected the 'error' to be 'invalid page id'. Got '%v'", m["error"])
		}
	})

	t.Run("fail - page not found", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&domain.Page{}, errors.New("unexpected error")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/pages/1", nil)
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.GetPage(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusNotFound, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if m["error"] != "page not found" {
			t.Errorf("Expected the 'error' to be 'page not found'. Got '%v'", m["error"])
		}
	})
}

func TestCtl_UpdatePage(t *testing.T) {
	mockSrv := new(service.MockPageService)
	mockPage := domain.Page{
		ID:        1,
		Content:   "page content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		BookID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockPage, nil).Once()
		mockSrv.On("Update", mock.Anything).Return(nil).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/pages/1", bytes.NewBuffer(jsonStr))
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.UpdatePage(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		if _, ok := m["result"]; !ok {
			t.Errorf("Expected `result` to exist. Got '%v'", m)
			return
		}
		result := m["result"].(map[string]interface{})
		if result["content"] != "test" {
			t.Errorf("Expected 'content' to be 'test'. Got '%v'", m["content"])
		}
		if result["bookID"] != 1.0 {
			t.Errorf("Expected `bookID` to be '1'. Got '%v'", m["bookID"])
		}
	})

	t.Run("fail - save error", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockPage, nil).Once()
		mockSrv.On("Update", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"title":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/book/1", bytes.NewBuffer(jsonStr))
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.UpdatePage(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
	})
}

func TestCtl_DeletePage(t *testing.T) {
	mockSrv := new(service.MockPageService)
	mockPage := domain.Page{
		ID:        1,
		Content:   "page content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		BookID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockPage, nil).Once()
		mockSrv.On("Delete", mock.Anything).Return(nil).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/pages/1", nil)
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.DeletePage(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
	})

	t.Run("fail - delete error", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockPage, nil).Once()
		mockSrv.On("Delete", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := controller.NewController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/pages/1", nil)
		vars := map[string]string{"id": "1"}
		r = mux.SetURLVars(r, vars)

		// run
		ctl.DeletePage(w, r)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
	})
}
