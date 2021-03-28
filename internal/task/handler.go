package task

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

	r.HandleFunc("/tasks", h.ListTasks).Methods("GET")
	r.HandleFunc("/task/{id:[0-9]+}", h.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}", h.GetTask).Methods("GET")
	r.HandleFunc("/task/{id:[0-9]+}", h.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{id:[0-9]+}", h.DeleteTask).Methods("DELETE")
}

// ListTasks returns a list of tasks
func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - list tasks")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// CreateTask returns a task
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - create task")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// GetTask returns a task
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get task")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// UpdateTask returns a task
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update task")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}

// DeleteTask returns a task
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete task")
	utils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}
