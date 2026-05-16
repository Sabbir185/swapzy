package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

func ConvertJsonToYaml(filePath string) (string, error) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var data interface{}

	err = yaml.Unmarshal(jsonData, &data)
	if err != nil {
		return "", err
	}

	yamlData, err := yaml.Marshal(&data)
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}
