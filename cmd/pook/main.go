package main

import (
	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/pook"
)

func main() {
	config := config.LoadEnv()
	app := pook.NewApp(config)
	app.Run()
}
