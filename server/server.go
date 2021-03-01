package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

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
	server.DB.AutoMigrate(Booking{})
	server.DB.AutoMigrate(Location{})

	// set secret
	if len(secret) < 0 {
		log.Fatal("Missing secret string")
	}
	server.Secret = secret

	// assign runtime address
	server.Addr = ":" + port

	// setup router
	server.Router = mux.NewRouter().StrictSlash(true)
	server.Router.Use(cors)
	server.Router.HandleFunc("/ping", server.pingHandler).Methods("GET")

	return server
}

// Run starts the server
func (s *Server) Run() {
	log.Println("Listening on address " + s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, s.Router))
}

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - ping")
	fmt.Fprintf(w, "Ready to serve requests")
}
