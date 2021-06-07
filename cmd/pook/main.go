package main

import (
	"github.com/chumnend/pook/internal/config"
	"github.com/chumnend/pook/internal/pook"
)

func main() {
	cfg := config.LoadEnv()
	app := pook.NewApp(cfg)
	app.Run()
}
