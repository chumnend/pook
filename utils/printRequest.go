package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// PrintRequest - prints request info to console
func PrintRequest(r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(dump))
}
