package controller

import (
	"net/http"

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
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not yet implemented"})
}

func (b *bookController) CreateBook(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not yet implemented"})
}

func (b *bookController) GetBook(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not yet implemented"})
}

func (b *bookController) UpdateBook(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not yet implemented"})
}

func (b *bookController) DeleteBook(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not yet implemented"})
}
