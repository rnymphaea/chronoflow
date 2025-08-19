package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type StorageConfig struct {
	DatabaseType string `env:"DATABASE_TYPE" envDefault:"postgres"`
	CacheType    string `env:"CACHE_TYPE" envDefault:"memory"`
}

type LoggerConfig struct {
	Type   string `env:"LOGGER_TYPE" envDefault:"zerolog"`
	Level  string `env:"LOGGER_LEVEL" envDefault:"info"`
	Pretty bool   `env:"LOGGER_PRETTY" envDefault:"false"`
}

type RetryConfig struct {
	InitialTimeout time.Duration `env:"RETRY_INITIAL_TIMEOUT" envDefault:"100ms"`
	Multiplier     float64       `env:"RETRY_MULTIPLIER" envDefault:"2.0"`
	Jitter         float64       `env:"RETRY_JITTER" envDefault:"0.2"`
	MaxAttempts    int           `env:"RETRY_MAX_ATTEMPTS" envDefault:"5"`
}

func LoadStorageConfig() (*StorageConfig, error) {
	var cfg StorageConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadLoggerConfig() (*LoggerConfig, error) {
	var cfg LoggerConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
