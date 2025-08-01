package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type ServerConfig struct {
	Address string `env:"ADDRESS" envDefault:":50051"`

	MaxConcurrentStreams uint32 `env:"MAX_CONCURRENT_STREAMS" envDefault:"100"`

	MaxConnectionIdle     time.Duration `env:"MAX_CONN_IDLE" envDefault:"30m"`
	MaxConnectionAge      time.Duration `env:"MAX_CONN_AGE" envDefault:"2h"`
	MaxConnectionAgeGrace time.Duration `env:"MAX_CONN_AGE_GRACE" envDefault:"3m"`

	Time    time.Duration `env:"TIME" envDefault:"2h"`
	Timeout time.Duration `env:"TIMEOUT" envDefault:"20s"`
}

func LoadServerConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
