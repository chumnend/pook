package pook

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chumnend/pook/internal/pook/config"
	"github.com/chumnend/pook/internal/pook/middleware"
	"github.com/chumnend/pook/internal/pook/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

// App struct declaration
type App struct {
	Config *config.Config
	Conn   *gorm.DB
	Router *mux.Router
}

// NewApp builds a new app instance with given configuration settings
func NewApp() *App {
	cfg := config.LoadEnv()

	db, err := gorm.Open("postgres", cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewPostgresRepository(db)
	userSrv := user.NewService(userRepo)
	userCtl := user.NewUserController(userSrv)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.Cors)

	// setup api subrouter
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/register", userCtl.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/login", userCtl.Login).Methods("POST", "OPTIONS")

	// serve react files on catchall handler
	spa := spaHandler{
		staticPath: cfg.StaticPath,
		indexPath:  cfg.IndexPath,
	}
	router.NotFoundHandler = spa

	return &App{
		Config: cfg,
		Conn:   db,
		Router: router,
	}
}

// Run starts the application
func (app *App) Run() {
	addr := ":" + app.Config.Port
	log.Println("Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
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
