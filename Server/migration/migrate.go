package migration

import (
	"server/initializers"
	"server/models"
)

//This is similar to Flyway migration for SpringBoot

func MigrateDatabases() {
	if initializers.DB == nil {
		//This will establish the connection with DB
		initializers.DB = initializers.EstablishDBConnection()
	}

	//Once the Connection is established, migration will be done.
	initializers.DB.AutoMigrate(&models.Customer{})
	initializers.DB.AutoMigrate(&models.Course{})
	initializers.DB.AutoMigrate(&models.Payment{})
}
