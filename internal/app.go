package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App struct declaration
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Init setups up the application
func (app *App) Init(connectionString string) {
	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router.HandleFunc("/", pingServer).Methods("GET")
}

// Run starts tthe web server
func (app *App) Run(addr string) {
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

// Handlers =====================================
func pingServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready to serve requests")
}
