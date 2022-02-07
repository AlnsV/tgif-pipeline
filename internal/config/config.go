package config

import (
	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

type Config struct {
	FTXAPIKey    string `env:"FTX_API_KEY"`
	FTXAPISecret string `env:"FTX_API_SECRET"`

	InfluxAddress string `env:"INFLUX_ADDRESS"`
}

func New() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, errors.Wrap(err, "error with initializing config")
	}

	return &config, nil
}
