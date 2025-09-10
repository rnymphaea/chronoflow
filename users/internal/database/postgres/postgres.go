package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rnymphaea/chronoflow/users/internal/config"
	"github.com/rnymphaea/chronoflow/users/internal/logger"
	"github.com/rnymphaea/chronoflow/users/internal/util"
)

type PostgresDB struct {
	pool           *pgxpool.Pool
	requestTimeout time.Duration
	retryCfg       config.RetryConfig

	log logger.Logger
}

func New(cfg *config.PostgresConfig, log logger.Logger) (*PostgresDB, error) {
	log = log.Component("postgres")

	var p PostgresDB
	p.log = log
	p.retryCfg = cfg.RetryCfg
	p.requestTimeout = cfg.RequestTimeout

	p.log.Debug("creating new postgres pool")

	dsn := fmt.Sprintf("postgres://%s:***@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	poolCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	poolCfg.MaxConns = cfg.PoolMaxConns
	poolCfg.MinConns = cfg.PoolMinConns
	poolCfg.MaxConnLifetime = cfg.PoolMaxConnLifetime
	poolCfg.MaxConnIdleTime = cfg.PoolMaxConnIdleTime
	poolCfg.HealthCheckPeriod = cfg.PoolHealthCheckPeriod

	log.Debugf("trying to create pool", map[string]interface{}{
		"dsn":          dsn,
		"retry_config": p.retryCfg,
	})

	for i := 0; i < p.retryCfg.MaxAttempts; i++ {
		ctx, cancel := context.WithTimeout(context.TODO(), p.requestTimeout)

		pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
		cancel()
		if err == nil {
			ctxPing, cancel := context.WithTimeout(context.TODO(), p.requestTimeout)
			defer cancel()
			if err = pool.Ping(ctxPing); err == nil {
				p.pool = pool
				p.log.Info("successfully connected to postgres")
				break
			}
		}

		delay := util.RetryDelay(p.retryCfg, i)

		p.log.Warnf("connection failed", map[string]interface{}{
			"attempt": i + 1,
			"err":     err,
		})

		time.Sleep(delay)
	}

	return &p, nil
}

func (p *PostgresDB) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func (p *PostgresDB) Close() {
	p.pool.Close()
}
