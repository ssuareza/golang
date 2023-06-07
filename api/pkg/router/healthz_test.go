package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	// svc := New()

	// set up testing server
	ts := httptest.NewServer(http.HandlerFunc(Healthz))
	defer ts.Close()

	// make request
	res, err := http.Get(ts.URL + "/healthz")
	if err != nil {
		t.Fatalf("unexpected error; got %v", err)
	}

	fmt.Println(res.StatusCode)
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d; got: %v", http.StatusOK, res.StatusCode)
	}
}
