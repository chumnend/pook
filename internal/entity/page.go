package entity

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Page represents a page of a book in the application
type Page struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	BookID    uint      `gorm:"not null" json:"bookID"`
}

// PageRepository is the contract between DB to application
type PageRepository interface {
	FindAllByBookID(id uint) ([]Page, error)
	FindByID(id uint) (*Page, error)
	Create(page *Page) error
	Update(page *Page) error
	Delete(page *Page) error
	Migrate() error
}

// PageService handles the business logic regarding Pages
type PageService interface {
	FindAllByBookID(uint) ([]Page, error)
	FindByID(uint) (*Page, error)
	Create(page *Page) error
	Update(page *Page) error
	Delete(*Page) error
	Validate(*Page) error
}

// PageController defines page handlers in the application
type PageController interface {
	ListPages(c *gin.Context)
	CreatePage(c *gin.Context)
	GetPage(c *gin.Context)
	UpdatePage(c *gin.Context)
	DeletePage(c *gin.Context)
}
