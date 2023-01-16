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

func TestPageController_ListPages(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSrv := new(pook_mock.MockPageService)
	mockPages := []entity.Page{
		{
			ID:        1,
			Content:   "page content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			BookID:    1,
		},
		{
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
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/pages?bookId=1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.ListPages(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["pages"]
		if !ok {
			t.Errorf("Expected `pages` to exist. Got '%v'", m)
			return
		}
		pages := value.([]interface{})
		if len(pages) != 2 {
			t.Errorf("Expected 'pages' to have length of 2. Got %v.", m["pages"])
		}
	})

	t.Run("fail - bookId not found", func(t *testing.T) {
		// setup
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/pages", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.ListPages(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		checkErrorMessage(t, w)
	})

	t.Run("fail - failed to to get book's pages", func(t *testing.T) {
		// setup
		mockSrv.On("FindAllByBookID", mock.AnythingOfType("uint")).Return([]entity.Page{}, errors.New("unable to access db")).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/pages?bookId=1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.ListPages(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
		checkErrorMessage(t, w)
	})
}

func TestBookController_CreatePage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSrv := new(pook_mock.MockPageService)

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Create", mock.Anything).Return(nil).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test", "bookID":"1"}`)
		r, _ := http.NewRequest("POST", "/v1/pages", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.CreatePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["page"]
		if !ok {
			t.Errorf("Expected `page` to exist. Got '%v'", m)
			return
		}
		page := value.(map[string]interface{})
		if page["content"] != "test" {
			t.Errorf("Expected 'content' to be 'test'. Got '%v'", page["content"])
		}
		if page["bookID"] != 1.0 {
			t.Errorf("Expected `bookID` to be '1'. Got '%v'", page["bookID"])
		}
	})

	t.Run("fail - bookId missing in request", func(t *testing.T) {
		// setup
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test"}`)
		r, _ := http.NewRequest("POST", "/v1/pages", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.CreatePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		checkErrorMessage(t, w)
	})

	t.Run("fail - page is invalid", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content": "test", "bookID": "1"}`)
		r, _ := http.NewRequest("POST", "/v1/pages", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.CreatePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		checkErrorMessage(t, w)
	})

	t.Run("fail - unable to create page", func(t *testing.T) {
		// setup
		mockSrv.On("Validate", mock.Anything).Return(nil).Once()
		mockSrv.On("Create", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content": "test", "bookID": "1"}`)
		r, _ := http.NewRequest("POST", "/v1/pages", bytes.NewBuffer(jsonStr))
		r.Header.Set("Content-Type", "application/json")
		c, _ := gin.CreateTestContext(w)
		c.Request = r

		// run
		ctl.CreatePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
		checkErrorMessage(t, w)
	})
}

func TestPageController_GetPage(t *testing.T) {
	mockSrv := new(pook_mock.MockPageService)
	mockPage := entity.Page{
		ID:        1,
		Content:   "page content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		BookID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockPage, nil).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/pages/1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.GetPage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["page"]
		if !ok {
			t.Errorf("Expected `page` to exist. Got '%v'", m)
			return
		}
		page := value.(map[string]interface{})
		assert.Equal(t, float64(mockPage.ID), page["id"]) // FixMe: Hacky comparison of uint
		assert.Equal(t, mockPage.Content, page["content"])
		assert.Equal(t, float64(mockPage.BookID), page["bookID"]) // FixMe: Hacky comparison of uint
	})

	t.Run("fail - invalid page id", func(t *testing.T) {
		// setup
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/pages/abc", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "abc",
			},
		}

		// run
		ctl.GetPage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		checkErrorMessage(t, w)
	})

	t.Run("fail - page not found", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&entity.Page{}, errors.New("unexpected error")).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/pages/1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.GetPage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusNotFound, w.Code)
		checkErrorMessage(t, w)
	})
}

func TestPageController_UpdatePage(t *testing.T) {
	mockSrv := new(pook_mock.MockPageService)
	mockPage := entity.Page{
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
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/pages/1", bytes.NewBuffer(jsonStr))
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.UpdatePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
		var m map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &m)
		value, ok := m["page"]
		if !ok {
			t.Errorf("Expected `page` to exist. Got '%v'", m)
			return
		}
		page := (value).(map[string]interface{})
		if page["content"] != "test" {
			t.Errorf("Expected 'content' to be 'test'. Got '%v'", m["content"])
		}
		if page["bookID"] != 1.0 {
			t.Errorf("Expected `bookID` to be '1'. Got '%v'", m["bookID"])
		}
	})

	t.Run("fail - invalid page id", func(t *testing.T) {
		// setup
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/pages/abc", bytes.NewBuffer(jsonStr))
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "abc",
			},
		}

		// run
		ctl.UpdatePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		checkErrorMessage(t, w)
	})

	t.Run("fail - invalid request body", func(t *testing.T) {
		// setup
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"contentless":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/pages/1", bytes.NewBuffer(jsonStr))
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.UpdatePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		checkErrorMessage(t, w)
	})

	t.Run("fail - unable to find page", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&entity.Page{}, errors.New("unexpected error")).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/pages/1", bytes.NewBuffer(jsonStr))
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		// run
		ctl.UpdatePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusNotFound, w.Code)
		checkErrorMessage(t, w)
	})

	t.Run("fail - unable to update", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockPage, nil).Once()
		mockSrv.On("Update", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		var jsonStr = []byte(`{"content":"test"}`)
		r, _ := http.NewRequest("PUT", "/v1/pages/1", bytes.NewBuffer(jsonStr))
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.UpdatePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
		checkErrorMessage(t, w)
	})
}

func TestPageController_DeletePage(t *testing.T) {
	mockSrv := new(pook_mock.MockPageService)
	mockPage := entity.Page{
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
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/pages/1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.DeletePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusOK, w.Code)
	})

	t.Run("fail - invalid page id", func(t *testing.T) {
		// setup
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/pages/abc", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "abc",
			},
		}

		// run
		ctl.DeletePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusBadRequest, w.Code)
		checkErrorMessage(t, w)
	})

	t.Run("fail - unable to find page", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&entity.Page{}, errors.New("unexpected error")).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/pages/1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.DeletePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusNotFound, w.Code)
		checkErrorMessage(t, w)
	})

	t.Run("fail - unable to update", func(t *testing.T) {
		// setup
		mockSrv.On("FindByID", mock.AnythingOfType("uint")).Return(&mockPage, nil).Once()
		mockSrv.On("Delete", mock.Anything).Return(errors.New("unexpected error")).Once()
		ctl := NewPageController(mockSrv)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/v1/pages/1", nil)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}

		// run
		ctl.DeletePage(c)

		// check
		mockSrv.AssertExpectations(t)
		checkResponseCode(t, http.StatusInternalServerError, w.Code)
		checkErrorMessage(t, w)
	})
}
