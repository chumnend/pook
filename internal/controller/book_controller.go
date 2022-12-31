package controller

import (
	"net/http"
	"strconv"

	"github.com/chumnend/pook/internal/entity"
	"github.com/gin-gonic/gin"
)

type bookController struct {
	srv entity.BookService
}

// NewBookController creates a BookController with given BookService
func NewBookController(srv entity.BookService) entity.BookController {
	return &bookController{srv: srv}
}

func (b *bookController) ListBooks(c *gin.Context) {
	var (
		books []entity.Book
		err   error
	)

	// check for userId in query
	userId := c.Query("userId")
	if userId != "" {
		uid64, _ := strconv.ParseUint(userId, 10, 64)
		books, err = b.srv.FindAllByUserID(uint(uid64))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
			return
		}
	} else {
		books, err = b.srv.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	UserID string `json:"userID" binding:"required"`
}

func (b *bookController) CreateBook(c *gin.Context) {
	// validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create book entity
	var book entity.Book
	book.Title = input.Title
	userID, _ := strconv.Atoi(input.UserID)
	book.UserID = uint(userID)

	// validate the new book struct
	if err := b.srv.Validate(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing and/or invalid information"})
		return
	}

	// save the book struct
	if err := b.srv.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func (b *bookController) GetBook(c *gin.Context) {
	// get book id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	// retrieve book
	book, err := b.srv.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

type UpdateBookInput struct {
	Title string `json:"title" binding:"required"`
}

func (b *bookController) UpdateBook(c *gin.Context) {
	// get book id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	// get book updated book info
	type requestBody struct {
		Title string `json:"title"`
	}

	// validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// load the book to be modified
	book, err := b.srv.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book does not exist"})
		return
	}

	// modify book fields
	book.Title = input.Title

	// save book
	if err := b.srv.Save(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func (b *bookController) DeleteBook(c *gin.Context) {
	// get book id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	// retrieve book
	book, err := b.srv.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book does not exist"})
		return
	}

	// delete book
	if err := b.srv.Delete(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete successful"})
}
