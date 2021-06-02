package main

import (
	"github.com/chumnend/pook/internal/app"
	"github.com/chumnend/pook/internal/config"
)

func main() {
	config := config.GetEnv()

	app := app.New()
	app.Initialize(config)
	app.Run()
}
