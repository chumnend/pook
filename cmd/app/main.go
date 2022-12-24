package main

import (
	"log"

	"github.com/chumnend/pook/config"
	"github.com/chumnend/pook/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
