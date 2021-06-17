package book

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/pook/domain"
	"github.com/chumnend/pook/internal/pook/response"
)

type bookCtl struct {
	srv domain.BookService
}

// NewController creates a BookController with given BookService
func NewController(srv domain.BookService) domain.BookController {
	return &bookCtl{srv: srv}
}

func (ctl *bookCtl) ListBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - list books")
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *bookCtl) CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - create book")
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *bookCtl) GetBook(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get book")
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *bookCtl) UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update book")
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *bookCtl) DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete book")
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}
