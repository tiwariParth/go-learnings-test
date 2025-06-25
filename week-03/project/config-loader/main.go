package main

import (
	"fmt"
	"os"
	"strings"

	"./config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: config-loader <config-file>")
		os.Exit(1)
	}
	
	configFile := os.Args[1]
	
	var loader config.ConfigLoader
	
	// Determine the config type based on file extension
	switch {
	case strings.HasSuffix(configFile, ".env"):
		loader = config.NewEnvFileConfig(configFile)
	case strings.HasSuffix(configFile, ".json"):
		loader = config.NewJSONConfig(configFile)
	case strings.HasSuffix(configFile, ".yaml"), strings.HasSuffix(configFile, ".yml"):
		loader = config.NewYAMLConfig(configFile)
	default:
		fmt.Printf("Unsupported file format: %s\n", configFile)
		fmt.Println("Supported formats: .env, .json, .yaml, .yml")
		os.Exit(1)
	}
	
	// Load the configuration
	err := loader.Load()
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}
	
	// TODO: Add code to display the loaded configuration
	// TODO: Add code to retrieve specific values from the configuration
	
	fmt.Println("Configuration loaded successfully!")
}
