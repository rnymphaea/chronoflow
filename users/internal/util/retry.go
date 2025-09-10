package util

import (
	"math"
	"math/rand"
	"time"

	"github.com/rnymphaea/chronoflow/users/internal/config"
)

func RetryDelay(cfg config.RetryConfig, attempt int) time.Duration {
	backoff := float64(cfg.InitialTimeout) * math.Pow(cfg.Multiplier, float64(attempt))
	jitter := backoff * cfg.Jitter * (rand.Float64()*2 - 1)

	delay := time.Duration(backoff + jitter)
	if delay < 0 {
		delay = cfg.InitialTimeout
	}

	return delay
}
