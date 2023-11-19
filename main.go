package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/Unity-Technologies/unity-gateway-y2y/controllers"
	"github.com/Unity-Technologies/unity-gateway-y2y/models"
)

func main() {
	configFile := flag.String("config", "config.yaml", "Path to the configuration file")
	outputFolder := flag.String("output", "output", "Path to the output folder")
	controllerFlags := flag.String("controllers", "rate_limiter", "Comma-separated list of controllers to run")
	outputFilename := flag.String("output-filename", "default.yaml", "Name of the output file")
	flag.Parse()

	config, err := models.LoadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	controllerNames := strings.Split(*controllerFlags, ",")

	for _, controllerName := range controllerNames {
		switch controllerName {
		case "rate_limiter":
			results, err := controllers.Render(config)
			if err != nil {
				panic(err)
			}

			err = controllers.WriteToFile(results, *outputFolder, *outputFilename)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Rendered YAML for %s controller saved to %s/%s\n", controllerName, *outputFolder, *outputFilename)

		default:
			fmt.Printf("Controller %s is not recognized\n", controllerName)
		}
	}
}
