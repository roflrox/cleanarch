package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/rs/zerolog/log"
)

func RunMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Cannot open db connection")
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "mysql", driver)
	if err != nil {
		log.Error().Err(err).Msg("Cannot apply migration")
		return err
	}
	err = m.Up()
	if err != nil {
		log.Error().Err(err).Msg("Cannot apply migration")
		return err
	}
	return nil
}
