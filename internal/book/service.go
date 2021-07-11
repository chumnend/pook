package book

import (
	"errors"

	"github.com/chumnend/pook/internal/domain"
)

type bookSrv struct {
	repo domain.BookRepository
}

// NewService returns a BookService utilizing provided BookRepository
func NewService(repo domain.BookRepository) domain.BookService {
	return &bookSrv{repo: repo}
}

func (srv *bookSrv) FindAll() ([]domain.Book, error) {
	books, err := srv.repo.FindAll()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (srv *bookSrv) FindAllByUserID(id uint) ([]domain.Book, error) {
	books, err := srv.repo.FindAllByUserID(id)
	if err != nil {
		return books, err
	}
	return books, nil
}

func (srv *bookSrv) FindByID(id uint) (*domain.Book, error) {
	book, err := srv.repo.FindByID(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (srv *bookSrv) Save(book *domain.Book) error {
	return srv.repo.Save(book)
}

func (srv *bookSrv) Delete(book *domain.Book) error {
	return srv.repo.Delete(book)
}

func (srv *bookSrv) Validate(book *domain.Book) error {
	if book == nil {
		return errors.New("Book is empty")
	}

	if book.Title == "" || book.UserID == 0 {
		return errors.New("Invalid Book")
	}
	return nil
}
