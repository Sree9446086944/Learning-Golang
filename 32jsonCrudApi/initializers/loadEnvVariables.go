package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() { // capital to export, make public
	//to load .env file
	err := godotenv.Load() // now port will be 3000, since given in .env file , loaded
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}
