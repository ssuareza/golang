package env

import (
	"fmt"
	"os"
)

// Env holds relevant env variables.
type Env struct {
	MemcachedURL string
}

// Init reads env vars.
func Init() (*Env, error) {
	env := Env{}
	var err error

	env.MemcachedURL, err = lookup("MEMCACHED_URL")
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
