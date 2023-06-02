package router

import (
	"time"

	"golang/zapier/pkg/api"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Router holds the logic to handle HTTP requests.
type Router struct {
	api api.API
}

// New returns a new Router.
func New() (Router, error) {
	return Router{
		api: api.New(e, l),
	}, nil
}

// Handler configures the service endpoints and returns the HTTP router.
func (r Router) Handler() *chi.Mux {
	svc := chi.NewRouter()

	// initialize go-chi middlewares
	svc.Use(middleware.RealIP)
	svc.Use(middleware.Recoverer)
	svc.Use(middleware.RedirectSlashes)
	svc.Use(render.SetContentType(render.ContentTypeJSON))
	svc.Use(middleware.Timeout(30 * time.Second))

	// define endpoints
	svc.Get("/healthz", r.api.Healthz)
	return svc
}
