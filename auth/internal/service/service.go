package service

import (
	"log"

	"github.com/rnymphaea/chronoflow/auth/internal/cache"
	"github.com/rnymphaea/chronoflow/auth/internal/config"
	"github.com/rnymphaea/chronoflow/auth/internal/logger"
)

type Service struct {
	Cache  cache.Cache
	Logger logger.Logger
}

func Run() {
	var s Service

	cfg, err := config.LoadServiceConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = s.setupComponents(cfg.Logger, cfg.Cache)
	if err != nil {
		log.Fatal(err)
	}

	s.Logger.Info("config loaded successfully", "cache", cfg.Cache.Type, "logger", cfg.Logger.Type)
}

func (s *Service) setupComponents(
	loggercfg config.LoggerConfig,
	cachecfg config.CacheConfig,
) error {
	err := s.setupLogger(loggercfg)
	if err != nil {
		return err
	}

	err = s.setupCache(cachecfg.Type, s.Logger)
	if err != nil {
		return err
	}

	return nil
}
