package pkg

import (
	"math"
	"math/rand"
	"time"
)

type RetryConfig struct {
	InitialTimeout time.Duration `env:"RETRY_INITIAL_TIMEOUT" envDefault:"100ms"`
	Multiplier     float64       `env:"RETRY_MULTIPLIER" envDefault:"2.0"`
	Jitter         float64       `env:"RETRY_JITTER" envDefault:"0.2"`
	MaxAttempts    int           `env:"RETRY_MAX_ATTEMPTS" envDefault:"5"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RetryDelay(cfg RetryConfig, attempt int) time.Duration {
	backoff := float64(cfg.InitialTimeout) * math.Pow(cfg.Multiplier, float64(attempt))
	jitter := backoff * cfg.Jitter * (rand.Float64()*2 - 1)

	delay := time.Duration(backoff + jitter)
	if delay < 0 {
		delay = cfg.InitialTimeout
	}

	return delay
}
