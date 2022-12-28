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
	return []entity.Book{}, errors.New("not yet implemented")
}

func (b *bookService) FindAllByUserID(uint) ([]entity.Book, error) {
	return []entity.Book{}, errors.New("not yet implemented")

}

func (b *bookService) FindByID(uint) (*entity.Book, error) {
	return &entity.Book{}, errors.New("not yet implemented")
}

func (b *bookService) Create(*entity.Book) error {
	return errors.New("not yet implemented")
}

func (b *bookService) Save(*entity.Book) error {
	return errors.New("not yet implemented")
}

func (b *bookService) Delete(*entity.Book) error {
	return errors.New("not yet implemented")
}

func (b *bookService) Validate(*entity.Book) error {
	return errors.New("not yet implemented")
}
