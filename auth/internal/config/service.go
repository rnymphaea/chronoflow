package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type ServiceConfig struct {
	Server ServerConfig
	Cache  CacheConfig
	Logger LoggerConfig
	Retry  RetryConfig
	JWT    JWTConfig
}

type CacheConfig struct {
	Type string `env:"CACHE_TYPE" envDefault:"memory"`
}

type LoggerConfig struct {
	Type   string `env:"LOGGER_TYPE"   envDefault:"zerolog"`
	Level  string `env:"LOGGER_LEVEL"  envDefault:"info"`
	Pretty bool   `env:"LOGGER_PRETTY" envDefault:"false"`
}

type RetryConfig struct {
	InitialTimeout time.Duration `env:"RETRY_INITIAL_TIMEOUT" envDefault:"100ms"`
	Multiplier     float64       `env:"RETRY_MULTIPLIER"      envDefault:"2.0"`
	Jitter         float64       `env:"RETRY_JITTER"          envDefault:"0.2"`
	MaxAttempts    int           `env:"RETRY_MAX_ATTEMPTS"    envDefault:"5"`
}

type JWTConfig struct {
	AccessSecret  string `env:"JWT_ACCESS_SECRET,file,required"`
	RefreshSecret string `env:"JWT_REFRESH_SECRET,file,required"`

	Issuer string `env:"JWT_ISSUER" envDefault:"auth"`

	AccessTTL  time.Duration `env:"JWT_ACCESS_TTL"  envDefault:"15m"`
	RefreshTTL time.Duration `env:"JWT_REFRESH_TTL" envDefault:"720h"`
}

func LoadServiceConfig() (*ServiceConfig, error) {
	var cfg ServiceConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
