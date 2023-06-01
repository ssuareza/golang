package main

import (
	// "log"
	"fmt"
	"net/http"
	"time"

	"github.com/ansel1/merry"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

const (
	serverAddr = "127.0.0.1:8000"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/health", HealthHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    serverAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Info().Str("addr", serverAddr).Msg("Server started")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(merry.Wrap(err).WithUserMessage("error starting server"))
	}
}

// HomeHandler is the home path response
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	logAccess(r)
}

// HealthHandler returns the health status of our application
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello\n")
	logAccess(r)
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

// logAccess logs the request
func logAccess(r *http.Request) {
	log.Info().
		Str("uri", r.RequestURI).
		Str("method", r.Method).
		Str("src", GetIP(r)).
		Msg("")
}
