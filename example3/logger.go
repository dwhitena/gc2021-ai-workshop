package main

import (
	"log"
	"net/http"
	"time"
)

// Logger receives a handler, wraps the passed handler with
// logging and timing.
func Logger(inner http.Handler, name string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})

}
