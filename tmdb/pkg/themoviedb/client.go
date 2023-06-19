package themoviedb

import (
	"net/http"
	"time"
)

// Client contains the client configuration
type Client struct {
	URL    string
	APIKey string
	Token  string
	Client *http.Client
}

// NewClient creates a new client
func NewClient(apiKey string) *Client {
	return &Client{
		URL:    APIURL,
		APIKey: apiKey,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
