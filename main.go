package main

import (
	"log"

	"github.com/connor-cahill/dropletAutomationCLI/cmd"
	"github.com/joho/godotenv"
)

// Init function will load in env vars from .env file
func init() {
	// Setup env vars from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error setting up env vars: ", err)
	}
}

func main() {

	// Execute the CLI package
	// and mount the root command
	// TODO: how can I pass the client obj to the CLI?
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatalln("Error starting the CLI: ", err)
	}
}
