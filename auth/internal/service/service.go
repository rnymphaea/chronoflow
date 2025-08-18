package service

import (
	"log"

	"github.com/rnymphaea/chronoflow/auth/internal/cache"
	"github.com/rnymphaea/chronoflow/auth/internal/config"
	"github.com/rnymphaea/chronoflow/auth/internal/database"
	"github.com/rnymphaea/chronoflow/auth/internal/logger"
)

type Service struct {
	Database database.Database
	Cache    cache.Cache
	Logger   logger.Logger
}

func Run() {
	var s Service

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

func (s *Service) registerComponents(loggercfg *config.LoggerConfig, storagecfg *config.StorageConfig) error {
	err := s.registerLogger(loggercfg)
	if err != nil {
		return err
	}

	err = s.registerDatabase(storagecfg.DatabaseType, s.Logger)
	if err != nil {
		return err
	}

	err = s.registerCache(storagecfg.CacheType, s.Logger)
	if err != nil {
		return err
	}

	return nil
}
