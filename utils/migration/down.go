package migration

import (
	"errors"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/golang-migrate/migrate/v4"
)

func MigrateDown(dbString string) error {
	migrationLogger := logging.GetLogger("migration", "down")

	m, err := migrate.New("file://migrations", dbString)
	if err != nil {
		return err
	}

	if err := m.Down(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		migrationLogger.Println("no migration changes")
	} else {
		migrationLogger.Println("database migrated")
	}

	return nil
}
