package themoviedb

import (
	"net/http"
	"net/http/httptest"
)

// TestingHTTPServer starts a local HTTP server for testing
func TestingHTTPServer(response string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}))

	return server
}
