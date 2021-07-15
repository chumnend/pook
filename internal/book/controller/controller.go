package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/domain"
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

	var (
		books []domain.Book
		err   error
	)

	// get a user's books if request body passed
	if r.Body != nil {
		type requestBody struct {
			ID uint `json:"userID"`
		}
		var request requestBody
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			respondWithError(w, http.StatusBadRequest, "something went wrong")
			return
		}
		defer r.Body.Close()

		books, err = ctl.srv.FindAllByUserID(request.ID)
	} else {
		books, err = ctl.srv.FindAll()
	}

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "something went wrong")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string][]domain.Book{"books": books})
}

func (ctl *bookCtl) CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - create book")
	respondWithError(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *bookCtl) GetBook(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get book")
	respondWithError(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *bookCtl) UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update book")
	respondWithError(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *bookCtl) DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete book")
	respondWithError(w, http.StatusNotImplemented, "Not yet implemented")
}
