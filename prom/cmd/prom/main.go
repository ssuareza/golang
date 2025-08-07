package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Create a prometheus counter
var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func main() {
	// Register the counter.
	prometheus.MustRegister(pingCounter)

	// Handlers.
	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

// ping handler.
func ping(w http.ResponseWriter, req *http.Request) {
	// Increment the counter.
	pingCounter.Inc()
	fmt.Fprintf(w, "Hello World!")
}
