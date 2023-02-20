package migration

import (
	"log"
	"os"
	"server/initializers"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//This is similar to Flyway migration for SpringBoot

func MigrateDatabases() {
	if initializers.DB == nil {
		//This will establish the connection with DB
		initializers.DB = initializers.EstablishDBConnection()
	}

	m, err := migrate.New(
		os.Getenv("DB_MIGRATION_URL"),
		os.Getenv("DB_MIGRATION_SCRIPTS_PATH"))

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
