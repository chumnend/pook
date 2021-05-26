package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chumnend/pook/internal/board"
	"github.com/chumnend/pook/internal/task"
	"github.com/chumnend/pook/internal/user"
	"github.com/chumnend/pook/internal/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
	"github.com/joho/godotenv"
)

// App struct declaration
type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

// NewApp returns an initialize App struct
func NewApp(connectionURL string) *App {
	// setup database connection
	db, err := gorm.Open("postgres", connectionURL)
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
		Router: router,
		DB:     db,
	}
}

// Serve sets the App to listen on given address
func (s *App) Serve(addr string) {
	log.Println("Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
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

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("missing env: SECRET_KEY")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing env: DATABASE_URL")
	}

	a := NewApp(dbURL)
	a.Serve(":" + port)
}
