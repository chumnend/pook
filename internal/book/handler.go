package book

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Handler struct declaration
type Handler struct {
	DB *gorm.DB
}

// AttachHandler takes a router and adds routes to it
func AttachHandler(r *mux.Router, db *gorm.DB) {
	h := &Handler{DB: db}

	r.HandleFunc("/books", h.ListBooksByUserID).Methods("GET")
	r.HandleFunc("/book/{id:[0-9]+}", h.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id:[0-9]+}", h.GetBook).Methods("GET")
	r.HandleFunc("/book/{id:[0-9]+}", h.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id:[0-9]+}", h.DeleteBook).Methods("DELETE")
}

// ListBooksByUserID returns a list of Books
func (h *Handler) ListBooksByUserID(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - list books")

	query := r.URL.Query()
	uid := query.Get("uid")

	if uid != "" {
		// get all books of a user
		books, err := ListBooksByUserID(h.DB, uid)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"results": books})
	} else {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'uid' not found")
		return
	}
}

// CreateBook returns a Book
func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - create book")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// GetBook returns a Book
func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get book")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// UpdateBook returns a Book
func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update book")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// DeleteBook returns a Book
func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete book")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}
