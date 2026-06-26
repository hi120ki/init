package logger

import (
	"log/slog"
	"os"

	"github.com/hi120ki/init/golang/internal/env"
)

func NewLogger(environment env.Environment) *slog.Logger {
	level := slog.LevelDebug
	if environment == env.EnvironmentProduction {
		level = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	slog.SetDefault(logger)

	return logger.With("environment", environment)
}
