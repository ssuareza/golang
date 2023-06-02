package main

import (
	"api/pkg/router"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

var (
	port = 8888
)

func main() {
	// initialize logger
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// initialize router
	r := router.New(log)

	log.Info().Msg("starting server")
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
		log.Panic().Err(err).Msg("server stopped")
	}
}
