package pook

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/pook/board"
	"github.com/chumnend/pook/internal/pook/task"
	"github.com/chumnend/pook/internal/pook/user"
	"github.com/chumnend/pook/internal/response"
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

// NewApp builds a new app instance with given configuration settings
func NewApp(config *config.Config) *App {
	app := App{Config: config}
	app.migrateDB()
	app.setupRouter()

	return &app
}

// Run starts the application
func (app *App) Run() {
	addr := ":" + app.Config.Port
	log.Println("Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) migrateDB() {
	var err error

	// setup database connection
	app.DB, err = gorm.Open("postgres", app.Config.DB)
	if err != nil {
		log.Fatal(err)
	}

	// migrate models to db
	app.DB.AutoMigrate(user.User{})
	app.DB.AutoMigrate(board.Board{})
	app.DB.AutoMigrate(task.Task{})
}

func (app *App) setupRouter() {
	// create router
	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router.Use(cors)

	// setup api subrouter
	api := app.Router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{"message": "Ready to serve requests"})
	}).Methods("GET")

	// attach api routes
	user.AttachHandler(api, app.DB)
	board.AttachHandler(api, app.DB)
	task.AttachHandler(api, app.DB)

	// serve react files on catchall handler
	spa := spaHandler{
		staticPath: app.Config.StaticPath,
		indexPath:  app.Config.IndexPath,
	}
	app.Router.NotFoundHandler = spa
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
