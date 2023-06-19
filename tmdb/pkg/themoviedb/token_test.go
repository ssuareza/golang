package themoviedb

import (
	"testing"
)

const (
	TokenResponse = `{
		"success": true,
		"expires_at": "2019-12-20 10:12:45 UTC",
		"request_token": "47ddb13dd075e5018d737b8e871c4bdd9c0e4611"
	}`
)

func TestNewToken(t *testing.T) {
	server := TestingHTTPServer(TokenResponse)
	defer server.Close()

	client := NewClient("xxxxxxxxxxxxxxxxxxxxxx")
	client.URL = server.URL

	token, err := client.NewToken()
	if err != nil || token != "47ddb13dd075e5018d737b8e871c4bdd9c0e4611" {
		t.Errorf("Not able to get expected token")
	}
}
