package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func EstablishDBConnection() (DB *gorm.DB) {
	if DB != nil {
		fmt.Println("DB connection is already established. Using the existing connection")
		return DB
	}

	fmt.Println("DB connection is not established. Trying to establish DB Connection")
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error in establishing DB connection")
		log.Fatal(err)
		return nil
	}
	return DB
}
