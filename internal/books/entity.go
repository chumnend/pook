package books

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Book represents a book in the Books table
type Book struct {
	Id        uuid.UUID `db:"id" json:"id"`
	UserId    uuid.UUID `db:"user_id" json:"userId"`
	Title     string    `db:"title" json:"title"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

// BookRepository is the contract between DB to application
type BookRepository interface {
	FindAll() ([]Book, error)
	FindAllByUserID(uuid.UUID) ([]Book, error)
	FindByID(uuid.UUID) (*Book, error)
	Create(*Book) error
	Save(*Book) error
	Delete(*Book) error
}

// BookService handles the business logic regarding Books
type BookService interface {
	FindAll() ([]Book, error)
	FindAllByUserID(uuid.UUID) ([]Book, error)
	FindByID(uuid.UUID) (*Book, error)
	Create(*Book) error
	Save(*Book) error
	Delete(*Book) error
	Validate(*Book) error
}

// BookController defines book handlers in the application
type BookController interface {
	ListBooks()
	CreateBook(w http.ResponseWriter, req *http.Request)
	GetBook(w http.ResponseWriter, req *http.Request)
	UpdateBook(w http.ResponseWriter, req *http.Request)
	DeleteBook(w http.ResponseWriter, req *http.Request)
}
