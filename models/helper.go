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

func NewMap(elements ...interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(elements); i += 2 {
		key, value := elements[i].(string), elements[i+1]
		m[key] = value
	}
	return m
}
