package pook

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chumnend/pook/internal/book"
	"github.com/chumnend/pook/internal/middleware"
	"github.com/chumnend/pook/internal/page"
	"github.com/chumnend/pook/internal/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

// App struct declaration
type App struct {
	Config *Config
	DB     *gorm.DB
	Router *mux.Router
}

// NewApp builds a new app instance
func NewApp() *App {
	// load config
	cfg := NewConfig()

	// connect database
	db, err := gorm.Open("postgres", cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	// setup router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.Cors)

	// attach api
	setupAPI(router, db)

	// serve react files on catchall handler
	spa := spaHandler{
		staticPath: cfg.StaticPath,
		indexPath:  cfg.IndexPath,
	}
	router.NotFoundHandler = spa

	return &App{
		Config: cfg,
		DB:     db,
		Router: router,
	}
}

// Run starts the application
func (app *App) Run() {
	addr := ":" + app.Config.Port
	log.Println("Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func setupAPI(router *mux.Router, db *gorm.DB) {
	// initialize repositories
	userRepo := user.NewPostgresRepository(db)
	if err := userRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	bookRepo := book.NewPostgresRepository(db)
	if err := bookRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	pageRepo := page.NewPostgresRepository(db)
	if err := pageRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	// initialize services
	userSrv := user.NewService(userRepo)
	bookSrv := book.NewService(bookRepo)
	pageSrv := page.NewService(pageRepo)

	// initialize controllers
	userCtl := user.NewController(userSrv)
	bookCtl := book.NewController(bookSrv)
	pageCtl := page.NewController(pageSrv)

	// setup api subrouter
	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/register", userCtl.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/login", userCtl.Login).Methods("POST", "OPTIONS")

	api.HandleFunc("/books", bookCtl.ListBooks).Methods("GET")
	api.HandleFunc("/books", bookCtl.CreateBook).Methods("POST", "OPTIONS")
	api.HandleFunc("/book/{id:[0-9]+}", bookCtl.GetBook).Methods("GET")
	api.HandleFunc("/book/{id:[0-9]+}", bookCtl.UpdateBook).Methods("PUT")
	api.HandleFunc("/book/{id:[0-9]+}", bookCtl.DeleteBook).Methods("DELETE")

	api.HandleFunc("/pages", pageCtl.ListPages).Methods("GET")
	api.HandleFunc("/pages", pageCtl.CreatePage).Methods("POST", "OPTIONS")
	api.HandleFunc("/page/{id:[0-9]+}", pageCtl.GetPage).Methods("GET")
	api.HandleFunc("/page/{id:[0-9]+}", pageCtl.UpdatePage).Methods("PUT")
	api.HandleFunc("/page/{id:[0-9]+}", pageCtl.DeletePage).Methods("DELETE")
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
