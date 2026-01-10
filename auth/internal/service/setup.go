package service

import (
	"fmt"

	"github.com/rnymphaea/chronoflow/auth/internal/cache"
	"github.com/rnymphaea/chronoflow/auth/internal/config"
	"github.com/rnymphaea/chronoflow/auth/internal/logger"
	zerolog "github.com/rnymphaea/chronoflow/auth/internal/logger/zerolog"
)

/*
func (s *Service) setupServer(cfg *config.ServerConfig) {
	ka := keepalive.ServerParameters{
		MaxConnectionIdle:     cfg.MaxConnectionIdle,
		MaxConnectionAge:      cfg.MaxConnectionAge,
		MaxConnectionAgeGrace: cfg.MaxConnectionAgeGrace,
		Time:                  cfg.Time,
		Timeout:               cfg.Timeout,
	}

	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(cfg.MaxConcurrentStreams),
		grpc.KeepaliveParams(ka),
	)

	healthServer := health.NewServer()
	healthpb.registerHealthServer(grpcServer, healthServer)

	healthServer.SetServingStatus("auth", healthpb.HealthCheckResponse_SERVING)

	s.GRPCServer = grpcServer
	s.Health = healthServer

	s.Logger.Info("gRPC server and health service registered")
}
*/

func (s *Service) setupLogger(cfg config.LoggerConfig) error {
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

func (s *Service) setupCache(cacheType string, log logger.Logger) error {
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
