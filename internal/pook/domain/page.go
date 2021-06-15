package domain

import (
	"net/http"
	"time"
)

// Page represents a page of a book in the application
type Page struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	BookID uint `gorm:"not null" json:"book_id"`
}

// PageRepository is the contract between DB to application
type PageRepository interface {
	FindAllByBookID(uint) ([]Page, error)
	FindByID(uint) (*Page, error)
	Save(*Page) error
	Delete(*Page) error
	Migrate() error
}

// PageService handles the business logic regarding Pages
type PageService interface {
	FindAllByBookID(uint) ([]Page, error)
	FindByID(uint) (*Page, error)
	Save(*Page) error
	Delete(*Page) error
	Migrate() error
	Validate(*Page) error
}

// PageController defines page handlers in the application
type PageController interface {
	ListPages(w http.ResponseWriter, r *http.Request)
	CreatePage(w http.ResponseWriter, r *http.Request)
	GetPage(w http.ResponseWriter, r *http.Request)
	UpdatePage(w http.ResponseWriter, r *http.Request)
	DeletePage(w http.ResponseWriter, r *http.Request)
}
