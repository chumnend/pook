package handlers

import (
	"net/http"

	"github.com/chumnend/pook/internal/models"
	"github.com/chumnend/pook/internal/utils"
	"github.com/google/uuid"
)

func CreateBook(w http.ResponseWriter, req *http.Request) {
	type requestInput struct {
		UserID string `json:"userId"`
		Title  string `json:"title"`
	}

	var input requestInput
	if err := utils.ParseJSON(req, &input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if input.Title == "" || input.UserID == "" {
		http.Error(w, "all fields (title, userId) are required", http.StatusBadRequest)
		return
	}

	parsedUUID, err := uuid.Parse(input.UserID)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if err := models.CreateBook(parsedUUID, input.Title); err != nil {
		http.Error(w, "unable to create book", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"message": "book successfully created",
	}

	utils.SendJSON(w, response, http.StatusOK)
}

func GetAllBooks(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")
	if userID != "" {
		parsedUUID, err := uuid.Parse(userID)
		if err != nil {
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}

		books, err := models.GetBooksByUserID(parsedUUID)
		if err != nil {
			http.Error(w, "unable to get user's books", http.StatusInternalServerError)
			return
		}

		response := map[string]any{
			"books": books,
		}

		utils.SendJSON(w, response, http.StatusOK)
	}

	books, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, "unable to get all book", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"books": books,
	}

	utils.SendJSON(w, response, http.StatusOK)
}

func GetBook(w http.ResponseWriter, req *http.Request) {
	bookID := req.PathValue("book_id")

	parsedUUID, err := uuid.Parse(bookID)
	if err != nil {
		http.Error(w, "invalid book_id", http.StatusBadRequest)
		return
	}

	book, err := models.GetBookByID(parsedUUID)
	if err != nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	response := map[string]any{
		"book": book,
	}

	utils.SendJSON(w, response, http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	bookID := req.PathValue("book_id")

	parsedUUID, err := uuid.Parse(bookID)
	if err != nil {
		http.Error(w, "invalid book_id", http.StatusBadRequest)
		return
	}

	type requestInput struct {
		Title string `json:"title"`
	}

	var input requestInput
	if err := utils.ParseJSON(req, &input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if input.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	if err := models.UpdateBookByID(parsedUUID, input.Title); err != nil {
		http.Error(w, "unable to update book", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "book successfully updated",
	}

	utils.SendJSON(w, response, http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, req *http.Request) {
	bookID := req.PathValue("book_id")

	parsedUUID, err := uuid.Parse(bookID)
	if err != nil {
		http.Error(w, "invalid book_id", http.StatusBadRequest)
		return
	}

	if err := models.DeleteBookByID(parsedUUID); err != nil {
		http.Error(w, "unable to delete book", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "book successfully deleted",
	}

	utils.SendJSON(w, response, http.StatusOK)
}
