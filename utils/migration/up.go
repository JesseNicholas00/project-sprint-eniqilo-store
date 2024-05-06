package migration

import (
	"errors"

	"github.com/KerakTelor86/GoBoiler/utils/logging"
	"github.com/golang-migrate/migrate/v4"
)

func MigrateUp(dbString string) error {
	migrationLogger := logging.GetLogger("migration", "up")

	m, err := migrate.New("file://migrations", dbString)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		migrationLogger.Println("no migration changes")
	} else {
		migrationLogger.Println("database migrated")
	}

	return nil
}
