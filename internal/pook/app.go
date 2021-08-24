package pook

import (
	"log"
	"net/http"

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
	router := MakeRouter(cfg, db)

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
