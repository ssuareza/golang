package middleware

import (
	"time"

	"github.com/rs/zerolog"

	"net/http"

	"github.com/go-chi/chi/middleware"
)

// Logger is a middleware that logs requests
func Logger(log zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				log.Info().
					Str("proto", r.Proto).
					Str("path", r.URL.Path).
					Str("query", r.URL.Query().Encode()).
					Str("latency", time.Since(time.Now()).String()).
					Int("status", ww.Status()).
					Int("size", ww.BytesWritten()).
					Msg("request")
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
