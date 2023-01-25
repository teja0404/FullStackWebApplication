package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func EstablishConnection() (DB *gorm.DB) {
	fmt.Println("Trying to establish DB Connection")
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error in establishing Database connection")
		return nil
	} else if DB != nil {
		fmt.Println("DB connection established")
		return DB
	}

	return DB
}
