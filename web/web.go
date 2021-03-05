package web

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
	Addr   string
	Router *mux.Router
	DB     *gorm.DB
}

// NewApp creates and setups up App instance
func NewApp(dbURL string, port string) *App {
	var err error

	app := new(App)

	// connect database
	app.DB, err = gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// setup api routes
	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router.HandleFunc("/status", statusHandler)

	app.Addr = ":" + port

	return app
}

// Start makes the server listen on given port
func (a *App) Start() {
	log.Println("Listening on " + a.Addr)
	log.Fatal(http.ListenAndServe(a.Addr, a.Router))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready to serve requests")
}
