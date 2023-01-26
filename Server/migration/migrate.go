package migration

import (
	"server/initializers"
	"server/models"
)

//This is similar to Flyway migration for SpringBoot

func MigrateDatabases() {
	if initializers.DB == nil {
		initializers.LoadEnvVariables()
		initializers.DB = initializers.EstablishConnection()
	}

	initializers.DB.AutoMigrate(&models.Customer{})
	initializers.DB.AutoMigrate(&models.Course{})
}
