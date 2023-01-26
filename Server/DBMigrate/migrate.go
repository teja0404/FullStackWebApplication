package main

import (
	"server/initializers"
	"server/models"
)

//This is similar to Flyway migration for SpringBoot

func init() {
	initializers.LoadEnvVariables()
	initializers.EstablishConnection()
}

func main() {
	initializers.DB.AutoMigrate(&models.Student{})
}
