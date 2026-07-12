package main

import (
	"log"

	"github.com/hi120ki/init/golang/internal/env"
	"github.com/hi120ki/init/golang/internal/logger"
)

func main() {
	cfg, err := env.Load()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	logger := logger.NewLogger(cfg.Environment)

	logger.Info("application started", "port", cfg.Port)
}
