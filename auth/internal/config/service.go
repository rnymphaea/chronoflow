package config

import (
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
