package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	Version     string
	ServiceName string
	DB_STRING   string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	db_string := os.Getenv("DB_STRING")

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		DB_STRING:   db_string,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
