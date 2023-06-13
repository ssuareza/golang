package env

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("MEMCACHED_URL", "10.0.0.1:11211")

	env, err := Init()

	if err != nil {
		t.Errorf("error: %s", err)
	}

	if env.MemcachedURL != "10.0.0.1:11211" {
		t.Errorf("expected: %s, got %s", "10.0.0.1:11211", env.MemcachedURL)
	}
}
