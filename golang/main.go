package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIKey      string      `envconfig:"API_KEY" required:"true"`
	Environment Environment `envconfig:"ENVIRONMENT" default:"development"`
}

type Environment string

const (
	EnvironmentDevelopment Environment = "development"
	EnvironmentProduction  Environment = "production"
)

func (e Environment) LogLevel() slog.Level {
	if e == EnvironmentProduction {
		return slog.LevelInfo
	}
	return slog.LevelDebug
}

func (e Environment) IsValid() bool {
	return e == EnvironmentDevelopment || e == EnvironmentProduction
}

func newLogger(environment Environment) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: environment.LogLevel(),
	}))
}

func main() {
	const defaultEnvironment = EnvironmentDevelopment
	bootstrapLogger := newLogger(defaultEnvironment)
	slog.SetDefault(bootstrapLogger)

	_ = godotenv.Load()

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		slog.Error("Failed to process environment variables", "error", err)
		os.Exit(1)
	}

	if !config.Environment.IsValid() {
		slog.Error("Invalid environment value", "environment", config.Environment)
		os.Exit(1)
	}

	logger := bootstrapLogger
	if config.Environment != defaultEnvironment {
		logger = newLogger(config.Environment)
	}
	slog.SetDefault(logger)

	logger.Info("Application started", "api_key", config.APIKey, "environment", config.Environment)
}
