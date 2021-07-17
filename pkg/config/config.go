package config

import (
	"errors"
	"os"
)

var configLoaded bool = false
var configDatabaseURI string
var configDatabaseName string
var configAppEnv string = "prod"

// LoadConfig loads configuration from environment variables.
func LoadConfig() error {
	configDatabaseURI = os.Getenv("DB_URI")
	if len(configDatabaseURI) == 0 {
		return errors.New("DB_URI environment variable not set")
	}

	configDatabaseName = os.Getenv("DB_NAME")
	if len(configDatabaseName) == 0 {
		return errors.New("DB_NAME environment variable not set")
	}

	configAppEnv = os.Getenv(("APP_ENV"))
	if len(configAppEnv) == 0 {
		configAppEnv = "prod"
	}

	configLoaded = true
	return nil
}

// GetDatabaseURI returns the cached DB_URI from the environment.
func GetDatabaseURI() string {
	if !configLoaded {
		panic("Config not loaded")
	}
	return configDatabaseURI
}

// GetDatabaseName returns the cached DB_NAME from the environment.
func GetDatabaseName() string {
	if !configLoaded {
		panic("Config not loaded")
	}
	return configDatabaseName
}

// GetAppEnv returns the cached APP_ENV from the environment.
func GetAppEnv() string {
	return configAppEnv
}
