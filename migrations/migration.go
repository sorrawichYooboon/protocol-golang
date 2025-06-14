package migrations

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sorrawichYooboon/protocol-golang/config"
)

func RunMigrations(cfg *config.Config) {
	dbPath := "file://./migrations"
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseDBName)

	m, err := migrate.New(dbPath, databaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Database migrations applied successfully.")
}
