package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
