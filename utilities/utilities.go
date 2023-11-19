package utilities

import (
	"fmt"
	"os"

	"github.com/Unity-Technologies/unity-gateway-y2y/types"

	"gopkg.in/yaml.v2"
)

func NewMap(elements ...interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(elements); i += 2 {
		key, value := elements[i].(string), elements[i+1]
		m[key] = value
	}
	return m
}

func WriteToFile(data, folderPath, fileName string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			return err
		}
	}

	outputFilePath := fmt.Sprintf("%s/%s", folderPath, fileName)

	err := os.WriteFile(outputFilePath, []byte(data), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Data saved to %s\n", outputFilePath)
	return nil
}

func LoadConfig(filename string) (*types.Config, error) {
	configData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config types.Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
