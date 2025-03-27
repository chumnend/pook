package utils

import (
	"log"
	"net/http"
)

func Pong(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling /ping at %s\n", req.URL.Path)

	w.Write([]byte("Pong\n"))
}

func NotImplemented(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Not yet implemented", http.StatusNotImplemented)
}
