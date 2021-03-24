package pook

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chumnend/pook/internal/user/domain"
	"github.com/chumnend/pook/internal/user/entity"
	"github.com/chumnend/pook/internal/user/handler"
	"github.com/chumnend/pook/internal/user/repository"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

// Server struct declaration
type Server struct {
	Router *mux.Router
}

// NewServer returns an initialize Server struct
func NewServer(connectionURL string) *Server {
	db, err := gorm.Open("postgres", connectionURL)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserRepository(db)
	userEntity := entity.NewUserEntity(userRepo)

	router := mux.NewRouter().StrictSlash(true)
	handler.NewUserHandler(router, userEntity)

	// serve react files on catchall handler
	spa := spaHandler{
		staticPath: "web/build",
		indexPath:  "index.html",
	}
	router.NotFoundHandler = spa

	return &Server{
		Router: router,
	}
}

// Serve sets the Server to listen on given address
func (s *Server) Serve(addr string) {
	log.Println("Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
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
