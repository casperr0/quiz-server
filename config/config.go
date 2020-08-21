package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {

	loadConfig()
}

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
