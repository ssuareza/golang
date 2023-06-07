package env

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("API_USER", "user")
	os.Setenv("API_PASS", "pass")
	os.Setenv("PORT", "8888")
	os.Setenv("SERVICE_NAME", "api")

	env, err := Init()

	if err != nil {
		t.Errorf("error: %s", err)
	}

	if env.APIUser != "user" {
		t.Errorf("expected: %s, got %s", "user", env.APIUser)
	}

	if env.APIPass != "pass" {
		t.Errorf("expected: %s, got %s", "pass", env.APIPass)
	}

	if env.Port != "8888" {
		t.Errorf("expected: %s, got %s", "8888", env.Port)
	}

	if env.ServiceName != "api" {
		t.Errorf("expected: %s, got %s", "api", env.ServiceName)
	}
}
