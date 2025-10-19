package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIKey      string `envconfig:"API_KEY" required:"true"`
	Environment string `envconfig:"ENVIRONMENT" default:"development"`
}

func main() {
	environment := os.Getenv("ENVIRONMENT")
	logLevel := slog.LevelDebug
	if environment == "production" {
		logLevel = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

	_ = godotenv.Load()

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		slog.Error("Failed to process environment variables", "error", err)
		os.Exit(1)
	}

	slog.Info("Application started", "api_key", config.APIKey, "environment", config.Environment)
}
