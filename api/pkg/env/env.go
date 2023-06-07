package env

import (
	"fmt"
	"os"
)

// Env holds relevant env variables.
type Env struct {
	APIUser     string
	APIPass     string
	Port        string
	ServiceName string
}

// Init reads env vars.
func Init() (*Env, error) {
	env := Env{}
	var err error

	env.APIUser, err = lookup("API_USER")
	if err != nil {
		return &env, err
	}

	env.APIPass, err = lookup("API_PASS")
	if err != nil {
		return &env, err
	}

	env.Port, err = lookup("PORT")
	if err != nil {
		return &env, err
	}

	env.ServiceName, err = lookup("SERVICE_NAME")
	if err != nil {
		return &env, err
	}

	return &env, nil
}

// lookup helps verifying an env var exists.
func lookup(s string) (string, error) {
	value, ok := os.LookupEnv(s)
	if !ok {
		return "", fmt.Errorf("env var %s not set", s)
	}

	return value, nil
}
