package models

import (
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig(filename string) (*Config, error) {
	configData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
