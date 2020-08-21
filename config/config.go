package config

import (
	"log"

	"github.com/joho/godotenv"
)

// OfficerConfig contains all officer-related configs.
type OfficerConfig struct {
	DefaultRoles []string
}

// UsingConfig intergrate all config group.
type UsingConfig struct {
	Officer OfficerConfig
}

var (
	// Config act as a uniform interface.
	Config = UsingConfig{
		Officer: OfficerConfig{
			DefaultRoles: []string{"Admin", "Maintainer", "Guest"},
		},
	}
)

func init() {

	loadEnv()
}

func loadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
