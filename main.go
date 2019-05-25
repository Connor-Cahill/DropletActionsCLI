package main

import (
	"log"
	"fmt"
	"os"
	"path/filepath"

	"github.com/connor-cahill/dropletActionsCLI/cmd"
	"github.com/joho/godotenv"
	"github.com/connor-cahill/dropletActionsCLI/dropletauth"
	"github.com/boltdb/bolt"
	homedir "github.com/mitchellh/go-homedir"
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
    
    // find users home directory
    home, err := homedir.Dir()
    if err != nil {
        log.Fatalln("Error setting up DB: ", err)
    }

    // set path to db in users home dir
    dbPath := filepath.Join(home, "dropletActions.db")


    // if no DO API Key set ask for it
    if os.Getenv("DIGITAL_OCEAN_KEY") == "" {
        // sets as env var
        dropletauth.GetToken()
        fmt.Println("Token was set.")
    }
	// Execute the CLI package
	// and mount the root command
	// TODO: how can I pass the client obj to the CLI?
	err = cmd.RootCmd.Execute()
	if err != nil {
		log.Fatalln("Error starting the CLI: ", err)
	}

}
