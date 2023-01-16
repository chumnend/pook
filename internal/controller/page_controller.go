package controller

import (
	"net/http"
	"strconv"

	"github.com/chumnend/pook/internal/entity"
	"github.com/gin-gonic/gin"
)

type pageController struct {
	srv entity.PageService
}

// NewPageController creates a PageController with given PageService
func NewPageController(srv entity.PageService) entity.PageController {
	return &pageController{srv: srv}
}

func (ctl *pageController) ListPages(c *gin.Context) {
	// check for bookId from query
	bookId := c.Query("bookId")
	if bookId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookId not given"})
		return
	}

	// retrieve pages from db
	bookID64, _ := strconv.ParseUint(bookId, 10, 64)
	pages, err := ctl.srv.FindAllByBookID(uint(bookID64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pages": pages})
}

type CreatePageInput struct {
	Content string `json:"content" binding:"required"`
	BookID  string `json:"bookID" binding:"required"`
}

func (ctl *pageController) CreatePage(c *gin.Context) {
	// validate input
	var input CreatePageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create page entity
	var page entity.Page
	page.Content = input.Content
	bookID, _ := strconv.Atoi(input.BookID)
	page.BookID = uint(bookID)

	// validate the new page struct
	if err := ctl.srv.Validate(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing and/or invalid information"})
		return
	}

	// save the page struct
	if err := ctl.srv.Create(&page); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

func (ctl *pageController) GetPage(c *gin.Context) {
	// get page id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page id"})
		return
	}

	// retrieve page
	page, err := ctl.srv.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

type UpdatePageInput struct {
	Content string `json:"content" binding:"required"`
}

func (ctl *pageController) UpdatePage(c *gin.Context) {
	// get page id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page id"})
		return
	}

	// validate input
	var input UpdatePageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// load the page to be modified
	page, err := ctl.srv.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "page does not exist"})
		return
	}

	// modify page fields
	page.Content = input.Content

	// save page
	if err := ctl.srv.Update(page); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

func (ctl *pageController) DeletePage(c *gin.Context) {
	// get page id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page id"})
		return
	}

	// retrieve page
	page, err := ctl.srv.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	// delete page
	if err := ctl.srv.Delete(page); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": "page delete successfully"})
}
