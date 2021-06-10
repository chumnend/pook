package domain

import (
	"net/http"
	"time"
)

// Book represents a user's book in the application
type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	UserID uint `json:"user_id"`
}

// BookRepository is the contract between DB to application
type BookRepository interface {
	FindAll() ([]Book, error)
	FindAllByUserID(uint) ([]Book, error)
	FindBookByID(uint) (*Book, error)
	Save(*Book) error
	Delete(*Book) error
}

// BookService handles the business logic regarding Books
type BookService interface {
	BookRepository
}

// BookController defines book handlers in the application
type BookController interface {
	ListBooks(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}
