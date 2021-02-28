package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

var mutex = &sync.Mutex{}
var count = 0

// Server struct
type Server struct {
	Addr   string
	Router *mux.Router
	DB     *gorm.DB
	Secret string
}

// NewServer creates a new server struct
func NewServer(connectionString string, secret string, port string) *Server {
	server := new(Server)

	// connect database
	var err error
	server.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// migrate the schema
	server.DB.AutoMigrate(User{})

	// set secret
	if len(secret) < 0 {
		log.Fatal("Missing secret string")
	}
	server.Secret = secret

	// set the running address
	server.Addr = ":" + port

	return server
}

// Run starts the server
func (s *Server) Run() {
	log.Println("Listening on address " + s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}
