package main

import (
	"Server/initializers"
	"Server/models"
)

//This is similar to Flyway migration for SpringBoot

func init() {
	initializers.LoadEnvVariables()
	initializers.EstablishConnection()
}

func main() {
	initializers.DB.AutoMigrate(&models.Student{})
}
