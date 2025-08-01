package config

import (
	"github.com/caarlos0/env/v11"
)

type ServerConfig struct {
	GRPCAddress string `env:"GRPC_ADDRESS"`
}

func LoadServerConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
