package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TimeZone *time.Location `envconfig:"TIME_ZONE" default:"UTC"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to process env: %w", err)
	}
	return &cfg, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get core config: %w", err)
		panic(err)
	}
	return config
}
