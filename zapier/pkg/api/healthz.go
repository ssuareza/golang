package api

import (
	"fmt"
	"net/http"
)

// HealthHandler returns status of the server
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
