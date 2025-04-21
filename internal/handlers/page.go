package handlers

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/models"
	"github.com/chumnend/pook/internal/utils"
	"github.com/google/uuid"
)

func CreatePage(w http.ResponseWriter, req *http.Request) {
	log.Println("Request made to" + req.URL.Path)

	bookID := req.PathValue("book_id")

	type requestInput struct {
		ImageURL  string `json:"imageUrl"`
		Caption   string `json:"caption"`
		PageOrder int    `json:"pageOrder"`
	}

	var input requestInput
	if err := utils.ParseJSON(req, &input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if input.ImageURL == "" || input.Caption == "" || input.PageOrder < 0 {
		http.Error(w, "all fields (bookId, imageUrl, caption, pageOrder) are required", http.StatusBadRequest)
		return
	}

	parsedUUID, err := uuid.Parse(bookID)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	if err := models.CreatePage(parsedUUID, input.ImageURL, input.Caption, input.PageOrder); err != nil {
		log.Printf("failed to create page: %v", err)
		http.Error(w, "failed to create page", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "page successfully created",
	}

	utils.SendJSON(w, response, http.StatusOK)
}

func GetPages(w http.ResponseWriter, req *http.Request) {
	log.Println("Request made to" + req.URL.Path)

	bookID := req.PathValue("book_id")

	parsedUUID, err := uuid.Parse(bookID)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	pages, err := models.GetPagesByBookID(parsedUUID)
	if err != nil {
		log.Printf("failed to get book's page: %v", err)
		http.Error(w, "failed to get book's pages", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"pages": pages,
	}

	utils.SendJSON(w, response, http.StatusOK)
}

func GetPage(w http.ResponseWriter, req *http.Request) {
	log.Println("Request made to" + req.URL.Path)

	pageID := req.PathValue("page_id")

	parsedUUID, err := uuid.Parse(pageID)
	if err != nil {
		http.Error(w, "invalid page id", http.StatusBadRequest)
		return
	}

	page, err := models.GetPageByID(parsedUUID)
	if err != nil {
		http.Error(w, "page not found", http.StatusBadRequest)
		return
	}

	response := map[string]any{
		"page": page,
	}

	utils.SendJSON(w, response, http.StatusOK)
}

func UpdatePage(w http.ResponseWriter, req *http.Request) {
	log.Println("Request made to" + req.URL.Path)

	pageID := req.PathValue("page_id")

	type requestInput struct {
		ImageURL  string `json:"imageUrl"`
		Caption   string `json:"caption"`
		PageOrder int    `json:"pageOrder"`
	}

	var input requestInput
	if err := utils.ParseJSON(req, &input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if input.ImageURL == "" || input.Caption == "" || input.PageOrder < 0 {
		http.Error(w, "all fields (imageUrl, caption, pageOrder) are required", http.StatusBadRequest)
		return
	}

	parsedUUID, err := uuid.Parse(pageID)
	if err != nil {
		http.Error(w, "invalid page id", http.StatusBadRequest)
		return
	}

	if err := models.UpdatePage(parsedUUID, input.ImageURL, input.Caption, input.PageOrder); err != nil {
		log.Printf("failed to update page: %v", err)
		http.Error(w, "failed to update page", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "page successfully updated",
	}

	utils.SendJSON(w, response, http.StatusOK)
}

func DeletePage(w http.ResponseWriter, req *http.Request) {
	log.Println("Request made to" + req.URL.Path)

	pageID := req.PathValue("page_id")

	parsedUUID, err := uuid.Parse(pageID)
	if err != nil {
		http.Error(w, "invalid page id", http.StatusBadRequest)
		return
	}

	if err := models.DeletePage(parsedUUID); err != nil {
		log.Printf("failed to delete page: %v", err)
		http.Error(w, "failed to delete page", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "page successfully deleted",
	}

	utils.SendJSON(w, response, http.StatusOK)
}
