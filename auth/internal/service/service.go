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

	servercfg, err := config.LoadServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	loggercfg, err := config.LoadLoggerConfig()
	if err != nil {
		log.Fatal(err)
	}

	storagecfg, err := config.LoadStorageConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = s.registerComponents(loggercfg, storagecfg)
	if err != nil {
		log.Fatal(err)
	}

	s.Logger.Info("config loaded successfully")
}

func (s *Service) registerComponents(
	servercfg *config.ServerConfig,
	loggercfg *config.LoggerConfig,
	storagecfg *config.StorageConfig,
) error {
	err := s.registerServer(servercfg)
	if err != nil {
		return err
	}

	err := s.registerLogger(loggercfg)
	if err != nil {
		return err
	}

	err = s.registerCache(storagecfg.CacheType, s.Logger)
	if err != nil {
		return err
	}

	return nil
}
