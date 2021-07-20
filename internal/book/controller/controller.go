package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/chumnend/pook/internal/domain"
	"github.com/gorilla/mux"
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

	// create new book struct
	type DummyBook struct {
		Title  string `json:"title"`
		UserID string `json:"userID"`
	}
	var dummyBook DummyBook
	if err := json.NewDecoder(r.Body).Decode(&dummyBook); err != nil {
		respondWithError(w, http.StatusBadRequest, "something went wrong")
		return
	}
	defer r.Body.Close()

	var book domain.Book
	book.Title = dummyBook.Title
	id, _ := strconv.Atoi(dummyBook.UserID)
	book.UserID = uint(id)

	// validate the new book struct
	validateErr := ctl.srv.Validate(&book)
	if validateErr != nil {
		respondWithError(w, http.StatusBadRequest, "missing and/or invalid information")
		return
	}

	// save the book struct
	if err := ctl.srv.Save(&book); err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	// return success
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"result": book})
}

func (ctl *bookCtl) GetBook(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get book")

	// get book id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	// retrieve book
	book, err := ctl.srv.FindByID(uint(id))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "book not found")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{"result": book})
}

func (ctl *bookCtl) UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update book")
	respondWithError(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *bookCtl) DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete book")
	respondWithError(w, http.StatusNotImplemented, "Not yet implemented")
}
