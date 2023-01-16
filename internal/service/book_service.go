package service

import (
	"errors"

	"github.com/chumnend/pook/internal/entity"
)

type bookService struct {
	repo entity.BookRepository
}

// NewBookService returns a BookService utilizing provided BookRepository
func NewBookService(repo entity.BookRepository) entity.BookService {
	return &bookService{repo: repo}
}

func (b *bookService) FindAll() ([]entity.Book, error) {
	books, err := b.repo.FindAll()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (b *bookService) FindAllByUserID(id uint) ([]entity.Book, error) {
	books, err := b.repo.FindAllByUserID(id)
	if err != nil {
		return books, err
	}
	return books, nil
}

func (b *bookService) FindByID(id uint) (*entity.Book, error) {
	book, err := b.repo.FindByID(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (b *bookService) Create(book *entity.Book) error {
	return b.repo.Create(book)
}

func (b *bookService) Save(book *entity.Book) error {
	return b.repo.Save(book)
}

func (b *bookService) Delete(book *entity.Book) error {
	return b.repo.Delete(book)
}

func (b *bookService) Validate(book *entity.Book) error {
	if book == nil {
		return errors.New("book is empty")
	}

	if book.Title == "" || book.UserID == 0 {
		return errors.New("invalid book")
	}
	return nil
}
