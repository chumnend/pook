package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/chumnend/pook/internal/api"
	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/db"
)

func main() {
	config.Init()
	db.Init()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down server...")
		if db.DB != nil {
			db.DB.Close()
		}
		os.Exit(0)
	}()

	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	log.Println("Starting server on port", config.Env.PORT)
	log.Fatal(http.ListenAndServe(":"+config.Env.PORT, mux))
}
