// Package api represents an HTTP API.
package api

// API holds API routes dependencies.
type API struct {
	env *env.Env
}

// New returns a new API.
func New(e *env.Env, l *log.Log) API {
	return API{
		env: e,
		log: l,
	}
}
