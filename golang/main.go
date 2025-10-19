package main

// Test comment for GitHub Actions workflow
import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIKey string `envconfig:"API_KEY" required:"true"`
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

	if err := godotenv.Load(); err != nil {
		slog.Warn("Failed to load .env file, using system environment variables", "error", err)
	} else {
		slog.Info("Loaded .env file")
	}

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		slog.Error("Failed to process environment variables", "error", err)
		os.Exit(1)
	}

	slog.Info("Application started", "api_key", config.APIKey)
}
