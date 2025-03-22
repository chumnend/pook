package handlers

import (
	"fmt"
	"net/http"
)

func Pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Pong\n")
}
