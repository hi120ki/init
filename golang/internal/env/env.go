package env

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Environment Environment `envconfig:"ENVIRONMENT" default:"development" validate:"required,oneof=development production"`
	Port        string      `envconfig:"PORT" default:"8080" validate:"required,numeric"`
}

type Environment string

const (
	EnvironmentDevelopment Environment = "development"
	EnvironmentProduction  Environment = "production"
)

func Load() (*Env, error) {
	_ = godotenv.Load()

	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, fmt.Errorf("failed to process envconfig: %w", err)
	}

	if err := validator.New().Struct(&env); err != nil {
		return nil, fmt.Errorf("failed to validate environment variables: %w", err)
	}

	return &env, nil
}
