package app

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chumnend/pook/internal/app/board"
	"github.com/chumnend/pook/internal/app/task"
	"github.com/chumnend/pook/internal/app/user"
	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

// App struct declaration
type App struct {
	Config *config.Config
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize sets up a web server application
func Initialize(config *config.Config) *App {
	// setup database connection
	db, err := gorm.Open("postgres", config.DB)
	if err != nil {
		log.Fatal(err)
	}

	// migrate models to db
	db.AutoMigrate(user.User{})
	db.AutoMigrate(board.Board{})
	db.AutoMigrate(task.Task{})

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(cors)

	// setup api subrouter
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Ready to serve requests"})
	}).Methods("GET")

	// attach api routes
	user.AttachHandler(api, db)
	board.AttachHandler(api, db)
	task.AttachHandler(api, db)

	// serve react files on catchall handler
	spa := spaHandler{
		staticPath: "web/build",
		indexPath:  "index.html",
	}
	router.NotFoundHandler = spa

	return &App{
		Config: config,
		Router: router,
		DB:     db,
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
