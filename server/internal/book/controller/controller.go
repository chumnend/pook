package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/chumnend/pook/server/internal/domain"
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

	// check for uid in query
	query := r.URL.Query()
	uid := query.Get("userId")
	if uid != "" {
		uid64, _ := strconv.ParseUint(uid, 10, 64)
		books, err = ctl.srv.FindAllByUserID(uint(uid64))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "something went wrong")
			return
		}
	} else {
		books, err = ctl.srv.FindAll()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "something went wrong")
			return
		}
	}

	respondWithJSON(w, http.StatusOK, map[string][]domain.Book{"books": books})
}

func (ctl *bookCtl) CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - create book")

	// create new book struct
	type requestBody struct {
		Title  string `json:"title"`
		UserID string `json:"userID"`
	}
	var request requestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "something went wrong")
		return
	}
	defer r.Body.Close()

	var book domain.Book
	book.Title = request.Title
	userID, _ := strconv.Atoi(request.UserID)
	book.UserID = uint(userID)

	// validate the new book struct
	if err := ctl.srv.Validate(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "missing and/or invalid information")
		return
	}

	// save the book struct
	if err := ctl.srv.Create(&book); err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{"book": book})
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

	respondWithJSON(w, http.StatusOK, map[string]interface{}{"book": book})
}

func (ctl *bookCtl) UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update book")

	// get book id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	// get book updated book info
	type requestBody struct {
		Title string `json:"title"`
	}
	var request requestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "something went wrong")
		return
	}
	defer r.Body.Close()

	// load the book to be modified
	book, err := ctl.srv.FindByID(uint(id))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "book not found")
		return
	}

	// modify book fields
	book.Title = request.Title

	// save book
	if err := ctl.srv.Save(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"book": book})
}

func (ctl *bookCtl) DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete book")

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

	// delete book
	if err := ctl.srv.Delete(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"result": "book delete successfully"})
}
