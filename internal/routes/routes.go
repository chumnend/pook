package routes

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/middleware"
	"github.com/chumnend/pook/internal/routes/book"
	"github.com/chumnend/pook/internal/routes/page"
	"github.com/chumnend/pook/internal/routes/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// NewRouter creates a new router struct to be used by the application
func NewRouter(cfg *config.Config, db *gorm.DB) *mux.Router {
	// create new router struct
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.Cors)

	// setup api routes
	api := router.PathPrefix("/v1").Subrouter()

	userCtl := user.MakeController(db)
	pageCtl := page.MakeController(db)
	bookCtl := book.MakeController(db)

	api.HandleFunc("/register", userCtl.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/login", userCtl.Login).Methods("POST", "OPTIONS")

	api.HandleFunc("/pages", pageCtl.ListPages).Methods("GET")
	api.HandleFunc("/pages", pageCtl.CreatePage).Methods("POST", "OPTIONS")
	api.HandleFunc("/pages/{id:[0-9]+}", pageCtl.GetPage).Methods("GET")
	api.HandleFunc("/pages/{id:[0-9]+}", pageCtl.UpdatePage).Methods("PUT")
	api.HandleFunc("/pages/{id:[0-9]+}", pageCtl.DeletePage).Methods("DELETE")

	api.HandleFunc("/books", bookCtl.ListBooks).Methods("GET")
	api.HandleFunc("/books", bookCtl.CreateBook).Methods("POST", "OPTIONS")
	api.HandleFunc("/books/{id:[0-9]+}", bookCtl.GetBook).Methods("GET")
	api.HandleFunc("/books/{id:[0-9]+}", bookCtl.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{id:[0-9]+}", bookCtl.DeleteBook).Methods("DELETE")

	// serve react files on catchall handler
	spa := spaHandler{
		staticPath: cfg.StaticPath,
		indexPath:  cfg.IndexPath,
	}
	router.NotFoundHandler = spa

	return router
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
