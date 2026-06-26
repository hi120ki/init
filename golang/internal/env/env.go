package env

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Environment Environment `envconfig:"ENVIRONMENT" default:"development"`
	Port        string      `envconfig:"PORT" default:"8080"`
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

	return &env, nil
}
