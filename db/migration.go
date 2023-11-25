package db

import (
	"app/cmd/config"
	"app/pkg/log"
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func Migration(config *config.Configuration, rollback bool) {
	// Create a new migration instance
	var dbURL string
	if config.Db.Vendor == "mysql" {
		dbURL = fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
			config.Db.Username, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Database)
	} else {
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
			config.Db.Username, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Database, config.Db.Schema)
	}

	m, err := migrate.New(
		"file://"+config.Dir.Migration,
		dbURL,
	)
	if err != nil {
		log.Fatal("Failed to create migration instance: " + err.Error())
	}

	// Run the specified migration action
	if rollback {
		err = m.Down()
	} else {
		err = m.Up()
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed")
	}

	// Get the current migration version
	version, dirty, err := m.Version()
	if err != nil {
		log.Fatal("Failed to get migration version")
	}

	log.System(fmt.Sprintf("Current migration version: %s (dirty: %s)\n", version, dirty))
}
