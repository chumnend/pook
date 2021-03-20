package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

// Server struct declaration
type Server struct {
	Addr   string
	Router *mux.Router
	DB     *gorm.DB
}

// New creates and setups up a Server struct
func New() *Server {
	return &Server{}
}

// Initialize the server
func (s *Server) Initialize(dbURL string, port string) {
	var err error

	// setup address
	s.Addr = ":" + port

	// connect database
	s.DB, err = gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// migrate schema
	s.DB.AutoMigrate(&User{})

	// setup router
	s.Router = mux.NewRouter().StrictSlash(true)

	// api routes
	api := s.Router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/status", s.statusHandler)
	api.HandleFunc("/v1/users", s.listUsersHandler).Methods("GET")
	api.HandleFunc("/v1/user", s.createUserHandler).Methods("POST")
	api.HandleFunc("/v1/user/{id:[0-9]+}", s.getUserHandler).Methods("GET")
	api.HandleFunc("/v1/user/{id:[0-9]+}", s.updateUserHandler).Methods("PUT")
	api.HandleFunc("/v1/user/{id:[0-9]+}", s.deleteUserHandler).Methods("DELETE")

	// serve react files on catchall handler
	spa := spaHandler{
		staticPath: "react/build",
		indexPath:  "index.html",
	}
	s.Router.NotFoundHandler = spa
}

// Run makes the server listen on given port
func (s *Server) Run() {
	log.Println("Listening on " + s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, s.Router))
}

func (s *Server) statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready to serve requests")
}

func (s *Server) listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := listUsers(s.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := user.createUser(s.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user := &User{ID: uint(id)}
	if err = user.getUser(s.DB); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (s *Server) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()

	user.ID = uint(id)

	if err := user.updateUser(s.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (s *Server) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	user := User{ID: uint(id)}
	if err := user.deleteUser(s.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (spa spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend path with path to static directory
	path = filepath.Join(spa.staticPath, path)
	if _, err = os.Stat(path); os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(spa.staticPath, spa.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(spa.staticPath)).ServeHTTP(w, r)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
