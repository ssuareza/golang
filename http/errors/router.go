package main

import (
	"log"
	"net/http"
)

// HandlerFunc is a helper for ResponseRouter struct
type HandlerFunc func(*http.Response)

// ResponseRouter gets the http response
type ResponseRouter struct {
	Handlers        map[int]HandlerFunc // int is equal to StatusCode
	DefaultHandlers HandlerFunc         // default function if not function defined for that specific StatusCode
}

// NewRouter is the constructor of ResponseRouter
func NewRouter() *ResponseRouter {
	return &ResponseRouter{
		Handlers: make(map[int]HandlerFunc),
		DefaultHandlers: func(r *http.Response) {
			log.Fatal("Unhandled response: ", r.StatusCode)
		},
	}
}

// Register matchs the response status code with a function we define. Check init() in errors.go.
func (r *ResponseRouter) Register(status int, handler HandlerFunc) {
	r.Handlers[status] = handler
}

// Process just process the response and returns the right handler
func (r *ResponseRouter) Process(resp *http.Response) {
	f, ok := r.Handlers[resp.StatusCode]
	if !ok {
		r.DefaultHandlers(resp)
		return
	}

	f(resp)
}
