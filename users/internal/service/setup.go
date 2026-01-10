package service

import (
	"fmt"

	"github.com/rnymphaea/chronoflow/users/internal/cache"
	"github.com/rnymphaea/chronoflow/users/internal/config"
	"github.com/rnymphaea/chronoflow/users/internal/database"
	"github.com/rnymphaea/chronoflow/users/internal/database/postgres"
	"github.com/rnymphaea/chronoflow/users/internal/logger"
	zerolog "github.com/rnymphaea/chronoflow/users/internal/logger/zerolog"
)

func (s *Service) registerLogger(cfg *config.LoggerConfig) error {
	var (
		l   logger.Logger
		err error
	)

	switch cfg.Type {
	case "zerolog":
		l = zerolog.New(cfg)
	default:
		return fmt.Errorf("logger type [%s] is not supported", cfg.Type)
	}

	s.Logger = l
	return err
}

func (s *Service) registerDatabase(dbType string, log logger.Logger) error {
	var (
		db  database.Database
		err error
	)

	switch dbType {
	case "postgres":
		cfg, err := config.LoadPostgresConfig()
		if err != nil {
			return err
		}

		db, err = postgres.New(cfg, log)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("database type [%s] is not supported", dbType)
	}

	s.Database = db
	return err
}

func (s *Service) registerCache(cacheType string, log logger.Logger) error {
	var (
		c   cache.Cache
		err error
	)

	switch cacheType {
	case "redis":
		//		cfg, err := config.LoadRedisConfig()
		//		if err != nil {
		//			return err
		//		}
		//
		//		c, err = redis.New(cfg, log)
		//		if err != nil {
		//			return err
		//		}

	default:
		return fmt.Errorf("cache type [%s] is not supported", cacheType)
	}

	s.Cache = c
	return err
}
