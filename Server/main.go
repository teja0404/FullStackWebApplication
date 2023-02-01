package main

import (
	"server/controller"
	"server/initializers"
	"server/migration"
)

func init() {
	initializers.LoadEnvVariables()
	migration.MigrateDatabases()
}

func main() {
	controller.HandleRequests()
}
