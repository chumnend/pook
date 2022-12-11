package main

import (
	"fmt"
	"os"

	"github.com/chumnend/pook/app/pook-api/webserver"
)

func main() {
	srv, err := webserver.New()
	if err != nil {
		fmt.Printf("Unable to create server instance: %s\n", err)
		os.Exit(1)
	}

	srv.Start()
}
