package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/db"
	"github.com/chumnend/pook/internal/routes"
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

	// Setup API routes
	routes.RegisterRoutes(mux)

	// Serve React static files with fallback to index.html
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reactBuildPath := filepath.Join("web", "dist")
		filePath := filepath.Join(reactBuildPath, r.URL.Path)

		// Check if the file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(reactBuildPath, "index.html"))
			return
		}

		http.FileServer(http.Dir(reactBuildPath)).ServeHTTP(w, r)
	})

	log.Println("Starting server on port", config.Env.PORT)
	log.Fatal(http.ListenAndServe(":"+config.Env.PORT, mux))
}
