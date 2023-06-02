package middleware

import (
	"net/http"
)

// BasicAuth is a middleware protecting routes with Basic Auth
func BasicAuth(username, password string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

			u, p, ok := r.BasicAuth()

			if !ok {
				http.Error(w, "not authorized", http.StatusUnauthorized)
				return
			}

			if u != username || p != password {
				http.Error(w, "n ot authorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
