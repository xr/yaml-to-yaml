package controllers

import (
	"fmt"
	"os"
)

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
