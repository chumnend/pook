package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/chumnend/pook/internal/domain"
	"github.com/gorilla/mux"
)

type pageCtl struct {
	srv domain.PageService
}

// NewController creates a PageController with given PageService
func NewController(srv domain.PageService) domain.PageController {
	return &pageCtl{srv: srv}
}

func (ctl *pageCtl) ListPages(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - list pages")

	// check for bookId in query
	query := r.URL.Query()
	bookID := query.Get("bookId")
	if bookID == "" {
		respondWithError(w, http.StatusBadRequest, "invalid bookId query")
		return
	}

	// retrieve pages from db
	bookID64, _ := strconv.ParseUint(bookID, 10, 64)
	pages, err := ctl.srv.FindAllByBookID(uint(bookID64))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string][]domain.Page{"pages": pages})
}

func (ctl *pageCtl) CreatePage(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - create page")

	// create new page struct
	type requestBody struct {
		Content string `json:"content"`
		BookID  string `json:"bookID"`
	}
	var request requestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "something went wrong")
		return
	}
	defer r.Body.Close()

	var page domain.Page
	page.Content = request.Content
	bookID, _ := strconv.Atoi(request.BookID)
	page.BookID = uint(bookID)

	// validate page struct
	if err := ctl.srv.Validate(&page); err != nil {
		respondWithError(w, http.StatusBadRequest, "missing and/or invalid information")
		return
	}

	// save the page struct
	if err := ctl.srv.Create(&page); err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"result": page})
}

func (ctl *pageCtl) GetPage(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get page")

	// get page id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid page id")
		return
	}

	// retrieve page
	page, err := ctl.srv.FindByID(uint(id))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "page not found")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{"result": page})
}

func (ctl *pageCtl) UpdatePage(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update page")

	// get page id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid page id")
		return
	}

	// get page info to be updated
	type requestBody struct {
		Content string `json:"content"`
	}
	var request requestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "something went wrong")
		return
	}
	defer r.Body.Close()

	// retrieve page to be updated
	page, err := ctl.srv.FindByID(uint(id))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "page not found")
		return
	}

	// modify page fields
	page.Content = request.Content

	// update page
	if err := ctl.srv.Update(page); err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"result": page})
}

func (ctl *pageCtl) DeletePage(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete page")

	// get page id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid page id")
		return
	}

	// retrieve page
	page, err := ctl.srv.FindByID(uint(id))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "page not found")
		return
	}

	// delete page
	if err := ctl.srv.Delete(page); err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"result": "page delete successfully"})
}
