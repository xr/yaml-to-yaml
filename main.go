package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/Unity-Technologies/unity-gateway-y2y/builders/rate_limiter"
	"github.com/Unity-Technologies/unity-gateway-y2y/utilities"
)

func main() {
	configFile := flag.String("config", "config.yaml", "Path to the configuration file")
	outputFolder := flag.String("output", "output", "Path to the output folder")
	builderFlags := flag.String("builders", "rate_limiter", "Comma-separated list of builders to run")
	outputFilename := flag.String("output-filename", "default.yaml", "Name of the output file")
	flag.Parse()

	config, err := utilities.LoadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	builderNames := strings.Split(*builderFlags, ",")

	for _, builderName := range builderNames {
		switch builderName {
		case "rate_limiter":
			results, err := rate_limiter.Render(config)
			if err != nil {
				panic(err)
			}

			err = utilities.WriteToFile(results, *outputFolder, *outputFilename)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Rendered YAML for %s builder saved to %s/%s\n", builderName, *outputFolder, *outputFilename)

		default:
			fmt.Printf("Builder %s is not recognized\n", builderName)
		}
	}
}
