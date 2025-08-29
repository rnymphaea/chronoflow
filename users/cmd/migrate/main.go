package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/rnymphaea/chronoflow/users/internal/config"
)

func main() {
	storagecfg, err := config.LoadStorageConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err = runMigrations(storagecfg); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully")
}

func runMigrations(cfg *config.StorageConfig) error {
	switch cfg.DatabaseType {
	case "postgres":
		return migratePostgres()
	default:
		return fmt.Errorf("database type [%s] is not supported", cfg.DatabaseType)
	}
}

func migratePostgres() error {
	cfg, err := config.LoadPostgresConfig()
	if err != nil {
		return err
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	m, err := migrate.New(
		"file:///migrations/postgres",
		dbURL,
	)

	if err != nil {
		return fmt.Errorf("failed to init migrate: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %v", err)
	}

	return nil
}
