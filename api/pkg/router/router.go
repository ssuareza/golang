package router

import (
	md "api/pkg/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"

	"github.com/go-chi/render"

	"net/http"
	"time"
)

var (
	username = "admin"
	password = "admin"
)

// Router holds the logic to handle HTTP requests.
type Router http.Handler

// Handler configures the service endpoints and returns the HTTP router.
func New(l zerolog.Logger) Router {
	// initialize go-chi router
	r := chi.NewRouter()

	// initialize custom log middleware
	r.Use(md.Logger(l))

	// initialize go-chi middlewares
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(30 * time.Second))

	// define endpoints
	r.Get("/healthz", Healthz)

	// define endpoints with basic auth
	r.Group(func(r chi.Router) {
		r.Use(md.BasicAuth(username, password))
		r.Get("/auth", Healthz)
	})

	return r
}
