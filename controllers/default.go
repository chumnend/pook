package controllers

import (
	"fmt"
	"net/http"
)

// Ready - informs client of api status
func Ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready to serve requests")
}
