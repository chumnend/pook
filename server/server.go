package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
func New(dbURL string, port string) *Server {
	// connect database
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// setup router
	router := mux.NewRouter().StrictSlash(true)

	// api routes
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/status", statusHandler)

	// ui routes
	spa := spaHandler{staticPath: "ui/build", indexPath: "index.html"}
	router.NotFoundHandler = spa

	return &Server{
		Addr:   ":" + port,
		Router: router,
		DB:     db,
	}
}

// Start makes the server listen on given port
func (server *Server) Start() {
	log.Println("Listening on " + server.Addr)
	log.Fatal(http.ListenAndServe(server.Addr, server.Router))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready to serve requests")
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
