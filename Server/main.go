package main

import (
	"fmt"
	"server/controller"
	"server/initializers"
	"server/migration"
)

func init() {
	fmt.Println("I am first")
	initializers.LoadEnvVariables()
	migration.MigrateDatabases()
}

func main() {
	fmt.Println("I am second")
	controller.HandleRequests()
}
