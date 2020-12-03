package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnvVars() {
	err := godotenv.Load("configuration.env")
	if err != nil {
		log.Print("Error loading .env file")
	}
}
