package board

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

	r.HandleFunc("/boards", h.ListBoardsByUserID).Methods("GET")
	r.HandleFunc("/boards", h.CreateBoard).Methods("POST", "OPTIONS")
	r.HandleFunc("/board/{id:[0-9]+}", h.GetBoard).Methods("GET")
	r.HandleFunc("/board/{id:[0-9]+}", h.UpdateBoard).Methods("PUT")
	r.HandleFunc("/board/{id:[0-9]+}", h.DeleteBoard).Methods("DELETE")
}

// ListBoardsByUserID returns a list of Boards
func (h *Handler) ListBoardsByUserID(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - list boards")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// get all boards of a user
	boards, err := ListBoardsByUserID(h.DB, userid)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"results": boards})
}

// CreateBoard returns a Board
func (h *Handler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	log.Println("POST - create board")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// create new board struct
	var b Board
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// parse UserID
	parsedUserID, _ := strconv.ParseUint(userid, 10, 64)
	b.UserID = uint(parsedUserID)

	// call method to create user in DB
	if err := b.Create(h.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"result": b})
}

// GetBoard returns a Board
func (h *Handler) GetBoard(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - get board")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// get board id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid board ID")
		return
	}

	// retrieve board
	board := Board{ID: uint(id)}
	if err := board.Get(h.DB); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "board not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"result": board})
}

// UpdateBoard returns a Board
func (h *Handler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT - update board")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// get board id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid board ID")
		return
	}

	// create new board struct
	var board Board
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&board); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	// modify fields
	board.ID = uint(id)

	// save the user
	if err := board.Update(h.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to update board")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"result": board})
}

// DeleteBoard returns a Board
func (h *Handler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE - delete board")

	query := r.URL.Query()

	// check for userid in query
	userid := query.Get("userid")
	if userid == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "query 'userid' not found")
		return
	}

	// get board id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid board ID")
		return
	}

	// delete the board
	board := Board{ID: uint(id)}
	if err := board.Delete(h.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to update board")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"result": "board delete successfully"})
}
