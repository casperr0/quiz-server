package config

import (
	"log"

	"github.com/joho/godotenv"
)

// OfficerConfig contains all officer-related configs.
type OfficerConfig struct {
	DefaultRoles []string
}

// QuizConfig contains all quiz-related configs.
type QuizConfig struct {
	DefaultTags []string
}

// FVTConfig contains all functional verification test config.
type FVTConfig struct {
	Topic   string
	Section string
}

// UsingConfig intergrate all config group.
type UsingConfig struct {
	Officer OfficerConfig
	Quiz    QuizConfig
	FVT     FVTConfig
}

var (
	// Config act as a uniform interface.
	Config = UsingConfig{
		Officer: OfficerConfig{
			DefaultRoles: []string{"Admin", "Maintainer", "Guest"},
		},
		Quiz: QuizConfig{
			DefaultTags: []string{
				"Network", "Language", "Security", "Hardware", "Animation",
				"Game", "SysAdmin", "School", "CCNS", "Engineering", "Math",
				"Others",
			},
		},
		FVT: FVTConfig{
			Topic:   "\n\n====================[ %s ]====================\n\n",
			Section: "\n---------------[ %s ]---------------\n",
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
