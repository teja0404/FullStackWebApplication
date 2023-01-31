package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error in initializing variables")
	} else {
		fmt.Println("Env variables initialized properly")
	}

}
