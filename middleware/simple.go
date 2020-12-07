package middleware

import (
	"log"
	"net/http"
)

// SimpleMiddleware - test middleware
func SimpleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("this middleware does nothing")
		next.ServeHTTP(w, r)
	})
}
