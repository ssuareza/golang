package main

import (
	"api/pkg/env"
	"api/pkg/router"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	// initialize environment variables
	env, err := env.Init()
	if err != nil {
		panic(err)
	}

	// initialize logger
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// initialize router
	r := router.New(env, log)

	log.Info().Msg("starting server")
	if err := http.ListenAndServe(fmt.Sprintf(":%v", env.Port), r); err != nil {
		log.Panic().Err(err).Msg("server stopped")
	}
}
