package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	http.HandleFunc("/health", HealthHandler)
	http.ListenAndServe(":8090", nil)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	logAccess(r)
	fmt.Fprintf(w, "hello\n")
}

// logAccess logs the request
func logAccess(r *http.Request) {
	log.Info().
		Str("uri", r.RequestURI).
		Str("method", r.Method).
		Msg("")
}
