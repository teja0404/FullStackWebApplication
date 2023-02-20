package main

import (
	"server/controller"
	"server/initializers"
	"server/migration"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	initializers.LoadEnvVariables()
	migration.MigrateDatabases()

}

func main() {
	controller.HandleRequests()
}
