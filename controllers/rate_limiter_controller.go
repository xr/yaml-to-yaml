package controllers

import (
	"fmt"

	"github.com/Unity-Technologies/unity-gateway-y2y/models"
	"gopkg.in/yaml.v2"
)

func Render(config *models.Config) (string, error) {

	envoyFilter := models.NewEnvoyFilter(config)

	yamlData, err := yaml.Marshal(envoyFilter)
	if err != nil {
		fmt.Printf("Error marshaling to YAML: %v\n", err)
	}

	return string(yamlData), nil
}
