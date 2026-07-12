package logger

import (
	"context"
	"log/slog"
	"testing"

	"github.com/hi120ki/init/golang/internal/env"
)

func TestNewLoggerDevelopmentEnablesDebug(t *testing.T) {
	logger := NewLogger(env.EnvironmentDevelopment)

	if !logger.Enabled(context.Background(), slog.LevelDebug) {
		t.Error("development logger should enable debug level")
	}
}

func TestNewLoggerProductionSuppressesDebug(t *testing.T) {
	logger := NewLogger(env.EnvironmentProduction)

	if logger.Enabled(context.Background(), slog.LevelDebug) {
		t.Error("production logger should suppress debug level")
	}
	if !logger.Enabled(context.Background(), slog.LevelInfo) {
		t.Error("production logger should enable info level")
	}
}
