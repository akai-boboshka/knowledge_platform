package config

import (
	"awesomeProject/pkg/models"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func InitConfig() (*models.Config, error) {
	file, err := os.Open("./internal/config/config.yaml")
	if err != nil {
		log.Printf("Error opening config file: %v", err)
		return nil, err
	}

	var config *models.Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Printf("Error decoding config file: %v", err)
		return nil, err
	}

	return config, nil
}
