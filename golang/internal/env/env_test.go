package env

import (
	"os"
	"testing"
)

func unset(t *testing.T, key string) {
	t.Helper()
	t.Setenv(key, "")
	os.Unsetenv(key)
}

func TestLoadDefaults(t *testing.T) {
	unset(t, "ENVIRONMENT")
	unset(t, "PORT")

	env, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}
	if env.Environment != EnvironmentDevelopment {
		t.Errorf("Environment = %q, want %q", env.Environment, EnvironmentDevelopment)
	}
	if env.Port != "8080" {
		t.Errorf("Port = %q, want %q", env.Port, "8080")
	}
}

func TestLoadReadsEnvironmentVariables(t *testing.T) {
	t.Setenv("ENVIRONMENT", "production")
	t.Setenv("PORT", "3000")

	env, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}
	if env.Environment != EnvironmentProduction {
		t.Errorf("Environment = %q, want %q", env.Environment, EnvironmentProduction)
	}
	if env.Port != "3000" {
		t.Errorf("Port = %q, want %q", env.Port, "3000")
	}
}

func TestLoadRejectsUnknownEnvironment(t *testing.T) {
	t.Setenv("ENVIRONMENT", "staging")
	unset(t, "PORT")

	if _, err := Load(); err == nil {
		t.Fatal("Load() succeeded, want error")
	}
}

func TestLoadRejectsNonNumericPort(t *testing.T) {
	unset(t, "ENVIRONMENT")
	t.Setenv("PORT", "not-a-port")

	if _, err := Load(); err == nil {
		t.Fatal("Load() succeeded, want error")
	}
}
