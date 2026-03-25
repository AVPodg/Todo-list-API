package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort string        `envconfig:"SERVER_PORT" default:":8080"`
	DbDSN      string        `envconfig:"DB_DSN" required:"true"`
	Timeout    time.Duration `envconfig:"TIMEOUT" default:"5s"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("APP", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
