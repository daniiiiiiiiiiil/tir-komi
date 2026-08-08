package jwt

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type JWTConfig struct {
	Secret          string        `envconfig:"JWT_SECRET" required:"true"`
	AccessTokenTTL  time.Duration `envconfig:"JWT_ACCESS_TTL" default:"15m"`
	RefreshTokenTTL time.Duration `envconfig:"JWT_REFRESH_TTL" default:"168h"`
}

func NewJWTConfig() (JWTConfig, error) {
	var cfg JWTConfig
	err := envconfig.Process("", &cfg)
	return cfg, err
}

func NewMustJWtConfig() JWTConfig {
	config, err := NewJWTConfig()
	if err != nil {
		panic(err)
	}
	return config
}
