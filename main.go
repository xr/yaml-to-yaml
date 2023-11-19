package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/Unity-Technologies/unity-gateway-y2y/controllers"
	"github.com/Unity-Technologies/unity-gateway-y2y/models"
)

func main() {
	// Define command-line flags for config file, output folder, controllers, and output filename
	configFile := flag.String("config", "config.yaml", "Path to the configuration file")
	outputFolder := flag.String("output", "output", "Path to the output folder")
	controllerFlags := flag.String("controllers", "rate_limiter", "Comma-separated list of controllers to run")
	outputFilename := flag.String("output-filename", "default.yaml", "Name of the output file")
	flag.Parse()

	// Load and parse the YAML configuration from the specified file
	config, err := models.LoadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	// Split the controller flags into a slice
	controllerNames := strings.Split(*controllerFlags, ",")

	// Iterate over the selected controllers and execute them
	for _, controllerName := range controllerNames {
		switch controllerName {
		case "rate_limiter":
			// Render rate limiter actions
			renderedActions, err := controllers.RenderRateLimiterActions(config)
			if err != nil {
				panic(err)
			}

			// Call the utility function to write the data to the specified output folder and filename
			err = controllers.WriteToFile(renderedActions, *outputFolder, *outputFilename)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Rendered YAML for %s controller saved to %s/%s\n", controllerName, *outputFolder, *outputFilename)

		// Add cases for other controllers as needed
		default:
			fmt.Printf("Controller %s is not recognized\n", controllerName)
		}
	}
}
