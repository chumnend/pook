package task

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/chumnend/pook/internal/response"
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
	r.HandleFunc("/tasks", h.CreateTask).Methods("POST", "OPTIONS")
	r.HandleFunc("/task/{id:[0-9]+}", h.GetTask).Methods("GET")
	r.HandleFunc("/task/{id:[0-9]+}", h.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{id:[0-9]+}", h.DeleteTask).Methods("DELETE")
}

// ListTasks returns a list of tasks
func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - list tasks")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		response.Error(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// check for boardId in query
	boardid := query.Get("boardid")
	if boardid == "" {
		response.Error(w, http.StatusBadRequest, "query 'boardid' not found")
		return
	}

	// get all tasks
	tasks, err := ListTasksByBoardID(h.DB, boardid)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{"results": tasks})
}

// CreateTask returns a task
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - create task")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		response.Error(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// check for boardId in query
	boardid := query.Get("boardid")
	if boardid == "" {
		response.Error(w, http.StatusBadRequest, "query 'boardid' not found")
		return
	}

	// create new task struct
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// parse UserID
	parsedUserID, _ := strconv.ParseUint(userid, 10, 64)
	t.UserID = uint(parsedUserID)

	// pasr BoardID
	parsedBoardID, _ := strconv.ParseUint(boardid, 10, 64)
	t.BoardID = uint(parsedBoardID)

	// call method to create user in DB
	if err := t.Create(h.DB); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{"result": t})
}

// GetTask returns a task
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get task")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		response.Error(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// check for boardId in query
	boardid := query.Get("boardid")
	if boardid == "" {
		response.Error(w, http.StatusBadRequest, "query 'boardid' not found")
		return
	}

	// get task id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid task ID")
		return
	}

	// retrieve board
	task := Task{ID: uint(id)}
	if err := task.Get(h.DB); err != nil {
		response.Error(w, http.StatusNotFound, "task not found")
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{"result": task})
}

// UpdateTask returns a task
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update task")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		response.Error(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// check for boardid in query
	boardid := query.Get("boardid")
	if boardid == "" {
		response.Error(w, http.StatusBadRequest, "query 'boardid' not found")
		return
	}

	// get task id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid task ID")
		return
	}

	// create new task struct
	var task Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	// modify fields
	task.ID = uint(id)

	// save the user
	if err := task.Update(h.DB); err != nil {
		response.Error(w, http.StatusInternalServerError, "unable to update task")
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{"result": task})
}

// DeleteTask returns a task
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete task")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		response.Error(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// check for boardid in query
	boardid := query.Get("boardid")
	if boardid == "" {
		response.Error(w, http.StatusBadRequest, "query 'boardid' not found")
		return
	}

	// get task id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid task ID")
		return
	}

	// delete the board
	task := Task{ID: uint(id)}
	if err := task.Delete(h.DB); err != nil {
		response.Error(w, http.StatusInternalServerError, "unable to update task")
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{"result": "task delete successfully"})
}
