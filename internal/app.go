package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

// App struct declaration
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Init setups up the application
func (app *App) Init(connectionString string) {
	// connect database
	var err error
	app.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// migrate the schema
	app.DB.AutoMigrate(User{})

	// setup routes
	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router.Use(cors)
	app.Router.HandleFunc("/", pingServer).Methods("GET")
}

// Run starts tthe web server
func (app *App) Run(addr string) {
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

// Middleware ===================================
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		next.ServeHTTP(w, r)
	})
}

// Handlers =====================================
func pingServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready to serve requests")
}
