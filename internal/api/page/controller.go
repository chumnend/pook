package page

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/api/domain"
	"github.com/chumnend/pook/internal/api/response"
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
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *pageCtl) CreatePage(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - create page")
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *pageCtl) GetPage(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get page")
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *pageCtl) UpdatePage(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update page")
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}

func (ctl *pageCtl) DeletePage(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete page")
	response.Error(w, http.StatusNotImplemented, "Not yet implemented")
}
