package entity

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Book represents a user's book in the application
type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uint      `gorm:"not null" json:"userID"`

	Pages []Page `json:"pages"`
}

// BookRepository is the contract between DB to application
type BookRepository interface {
	FindAll() ([]Book, error)
	FindAllByUserID(uint) ([]Book, error)
	FindByID(uint) (*Book, error)
	Create(*Book) error
	Save(*Book) error
	Delete(*Book) error
	Migrate() error
}

// BookService handles the business logic regarding Books
type BookService interface {
	FindAll() ([]Book, error)
	FindAllByUserID(uint) ([]Book, error)
	FindByID(uint) (*Book, error)
	Create(*Book) error
	Save(*Book) error
	Delete(*Book) error
	Validate(*Book) error
}

// BookController defines book handlers in the application
type BookController interface {
	ListBooks(c *gin.Context)
	CreateBook(c *gin.Context)
	GetBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}
