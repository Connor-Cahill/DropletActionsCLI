package main

import (
	"log"
	"fmt"
	"os"

	"github.com/connor-cahill/dropletActionsCLI/cmd"
	"github.com/joho/godotenv"
	"github.com/connor-cahill/dropletActionsCLI/dropletauth"
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

    // if no DO API Key set ask for it
    if os.Getenv("DIGITAL_OCEAN_KEY") == "" {
        // sets as env var
        dropletauth.GetToken()
        fmt.Println("Token was set.")
    }
	// Execute the CLI package
	// and mount the root command
	// TODO: how can I pass the client obj to the CLI?
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatalln("Error starting the CLI: ", err)
	}

}
