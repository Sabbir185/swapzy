package utils

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppName string `yaml:"APP_NAME"`
	AppEnv  string `yaml:"APP_ENV"`
	AppPort int    `yaml:"APP_PORT"`

	DatabaseHost     string `yaml:"DATABASE_HOST"`
	DatabasePort     int    `yaml:"DATABASE_PORT"`
	DatabaseUser     string `yaml:"DATABASE_USER"`
	DatabasePassword string `yaml:"DATABASE_PASSWORD"`
	DatabaseName     string `yaml:"DATABASE_NAME"`

	JWTSecret string `yaml:"JWT_SECRET"`

	RedisHost string `yaml:"REDIS_HOST"`
	RedisPort int    `yaml:"REDIS_PORT"`

	APIKey string `yaml:"API_KEY"`

	LogLevel string `yaml:"LOG_LEVEL"`
}

func ConvertYamlToJson(filePath string) (string, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var data Config

	err = yaml.Unmarshal(src, &data)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(&data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
