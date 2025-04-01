package books

import (
	"errors"

	"github.com/google/uuid"
)

type bookService struct {
	repo BookRepository
}

// NewBookService returns a BookService utilizing provided BookRepository
func NewBookService(repo BookRepository) BookService {
	return &bookService{repo: repo}
}

func (b *bookService) FindAll() ([]Book, error) {
	books, err := b.repo.FindAll()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (b *bookService) FindAllByUserID(id uuid.UUID) ([]Book, error) {
	books, err := b.repo.FindAllByUserID(id)
	if err != nil {
		return books, err
	}
	return books, nil
}

func (b *bookService) FindByID(id uuid.UUID) (*Book, error) {
	book, err := b.repo.FindByID(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (b *bookService) Create(book *Book) error {
	return b.repo.Create(book)
}

func (b *bookService) Save(book *Book) error {
	return b.repo.Save(book)
}

func (b *bookService) Delete(book *Book) error {
	return b.repo.Delete(book)
}

func (b *bookService) Validate(book *Book) error {
	if book == nil {
		return errors.New("book is empty")
	}

	if book.Title == "" || book.UserId == uuid.Nil {
		return errors.New("invalid book")
	}
	return nil
}
