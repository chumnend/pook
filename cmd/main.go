package main

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/api"
	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/db"
)

func main() {
	config.Init()

	db.Init()

	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	log.Println("Starting server on port", config.Env.PORT)
	log.Fatal(http.ListenAndServe(":"+config.Env.PORT, mux))
}
