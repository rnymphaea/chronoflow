package util_test

import (
	"testing"
	"time"

	"github.com/rnymphaea/chronoflow/auth/internal/config"
	"github.com/rnymphaea/chronoflow/auth/internal/util"
)

func TestRetryDelay(t *testing.T) {
	cfg := config.RetryConfig{
		InitialTimeout: 100 * time.Millisecond,
		Multiplier:     2.0,
		Jitter:         0.2,
		MaxAttempts:    5,
	}

	tests := []struct {
		name     string
		attempt  int
		wantBase time.Duration
	}{
		{"first attempt", 0, 100 * time.Millisecond},
		{"second attempt", 1, 200 * time.Millisecond},
		{"third attempt", 2, 400 * time.Millisecond},
		{"fourth attempt", 3, 800 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := util.RetryDelay(cfg, tt.attempt)

			want := float64(tt.wantBase)

			lower := want * (1 - cfg.Jitter)
			upper := want * (1 + cfg.Jitter)

			if float64(got) < lower || float64(got) > upper {
				t.Errorf("got delay %v, want around %v (expected between %v and %v)",
					got, tt.wantBase, time.Duration(lower), time.Duration(upper))
			}
		})
	}
}

func TestRetryDelay_NegativeJitter(t *testing.T) {
	cfg := config.RetryConfig{
		InitialTimeout: 100 * time.Millisecond,
		Multiplier:     2.0,
		Jitter:         1.0,
		MaxAttempts:    5,
	}

	got := util.RetryDelay(cfg, 0)
	if got <= 0 {
		t.Errorf("expected delay > 0, got %v", got)
	}
}

func BenchmarkRetryDelay(b *testing.B) {
	cfg := config.RetryConfig{
		InitialTimeout: 50 * time.Millisecond,
		Multiplier:     1.5,
		Jitter:         0.3,
		MaxAttempts:    5,
	}

	for i := 0; i < b.N; i++ {
		_ = util.RetryDelay(cfg, i%cfg.MaxAttempts)
	}
}

func TestRetryDelay_CoversNegativeDelayBranch(t *testing.T) {
	cfg := config.RetryConfig{
		InitialTimeout: 100 * time.Millisecond,
		Multiplier:     2.0,
		Jitter:         10.0,
		MaxAttempts:    5,
	}

	attempt := 0

	for i := 0; i < 100; i++ {
		delay := util.RetryDelay(cfg, attempt)
		if delay <= 0 {
			t.Errorf("expected delay > 0, got %v", delay)
		}
	}
}
