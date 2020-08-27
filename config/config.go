package config

import (
	"log"

	"github.com/joho/godotenv"
)

// OfficerConfig contains all officer-related configs.
type OfficerConfig struct {
	DefaultRoles []string
}

// LoadConfig contains all external data loading config.
type LoadConfig struct {
	ProdDir string
	DevDir  string
}

// FVTConfig contains all functional verification test config.
type FVTConfig struct {
	Topic   string
	Section string
	Detail  string
}

// UsingConfig intergrate all config group.
type UsingConfig struct {
	Officer OfficerConfig
	Load    LoadConfig
	FVT     FVTConfig
}

var (
	// Config act as a uniform interface.
	Config = UsingConfig{
		Officer: OfficerConfig{
			DefaultRoles: []string{"Admin", "Maintainer", "Guest"},
		},
		Load: LoadConfig{
			ProdDir: "data/",
			DevDir:  "example_data/",
		},
		FVT: FVTConfig{
			Topic:   "\n\n# %s\n",
			Section: "\n\n## %s\n",
			Detail:  "\n- [ %s ]\n",
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
