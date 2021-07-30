package main

import "github.com/chumnend/pook/server/internal/pook"

func main() {
	app := pook.NewApp()
	app.Run()
}
