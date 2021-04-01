package book

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	r.HandleFunc("/books", h.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id:[0-9]+}", h.GetBook).Methods("GET")
	r.HandleFunc("/book/{id:[0-9]+}", h.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id:[0-9]+}", h.DeleteBook).Methods("DELETE")
}

// ListBooksByUserID returns a list of Books
func (h *Handler) ListBooksByUserID(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - list books")

	query := r.URL.Query()

	// check for uid in query
	if uid := query.Get("uid"); uid != "" {
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

	query := r.URL.Query()

	// check for uid in query
	if uid := query.Get("uid"); uid != "" {
		// create new user struct
		var b Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		// parse UserID
		parsedUserID, _ := strconv.ParseUint(uid, 10, 64)
		b.UserID = uint(parsedUserID)

		// call method to create user in DB
		if err := b.Create(h.DB); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"result": b})
	} else {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'uid' not found")
		return
	}
}

// GetBook returns a Book
func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get book")

	query := r.URL.Query()

	// check for uid in query
	if uid := query.Get("uid"); uid != "" {
		// get book id
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "invalid book ID")
			return
		}

		// retrieve book
		book := Book{ID: uint(id)}
		if err := book.Get(h.DB); err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "book not found")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"result": book})
	} else {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'uid' not found")
		return
	}
}

// UpdateBook returns a Book
func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update book")

	query := r.URL.Query()

	// check for uid in query
	if uid := query.Get("uid"); uid != "" {
		// get book id
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "invalid book ID")
			return
		}

		// create new book struct
		var book Book
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&book); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "invalid request payload")
			return
		}
		defer r.Body.Close()

		// modify fields
		book.ID = uint(id)

		// save the user
		if err := book.Update(h.DB); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "unable to update book")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"result": book})
	} else {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'uid' not found")
		return
	}
}

// DeleteBook returns a Book
func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete book")

	query := r.URL.Query()

	// check for uid in query
	if uid := query.Get("uid"); uid != "" {
		// get book id
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "invalid book ID")
			return
		}

		// delete the book
		book := Book{ID: uint(id)}
		if err := book.Delete(h.DB); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "unable to update book")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"result": "book delete successfully"})
	} else {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'uid' not found")
		return
	}
}
