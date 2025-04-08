package api

import (
	"log"
	"net/http"
)

func Pong(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling %s\n", req.URL.Path)

	w.Write([]byte("Pong\n"))
}
