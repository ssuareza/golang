package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	port = 8080
)

// write an api using go-chi
func main() {
	// set go-chi router
	r := chi.NewRouter()

	// handlers
	r.Get("/healthz", r.api.Healthz)

	log.Fatal(http.ListenAndServe(":"+port, r))

}
