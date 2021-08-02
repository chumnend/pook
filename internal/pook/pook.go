package pook

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/routes"
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

// NewApp builds a new app instance
func NewApp() *App {
	// load config
	cfg := config.NewConfig()

	// connect database
	db, err := gorm.Open("postgres", cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	// setup router
	router := routes.NewRouter(cfg, db)

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
