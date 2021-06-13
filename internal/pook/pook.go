package pook

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chumnend/pook/internal/pook/config"
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

// NewApp builds a new app instance
func NewApp() *App {
	// load config
	cfg := config.GetEnv()

	// connect database
	db, err := gorm.Open("postgres", cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	// setup user
	userRepo := user.NewPostgresRepository(db)
	if err := userRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	userSrv := user.NewService(userRepo)
	userCtl := user.NewController(userSrv)

	// setup router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(cors)

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

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Allow-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Endcoding, Content-Type, Content-Length, Authorization, X-CSRF-token")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
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
