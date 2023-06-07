package router

import (
	"net/http"

	"github.com/go-chi/render"
)

type healthResponse struct {
	Status string `json:"status"`
}

// Healthz HTTP handler is an HTTP probe for service status.
func Healthz(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{
		Status: "ok",
	}

	render.JSON(w, r, resp)
}
