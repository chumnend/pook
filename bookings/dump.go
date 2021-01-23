package bookings

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func dumpRequest(r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(dump))
}
