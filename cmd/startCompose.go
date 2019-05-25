package cmd

import (
	"fmt"
	"log"

	dropletactions "github.com/connor-cahill/dropletActionsCLI/dropletActions"
	"github.com/spf13/cobra"
)

// StartComposeCmd will ssh into DO Droplet and run docker-compose up -d on compose file
// takes in DO Droplet public IP and path to compose file as arguments
var StartComposeCmd = &cobra.Command{
	Use:   "startCompose",
    Short: "Will ssh into DO Droplet and docker-compose up repo. Must pass following args IN ORDER:\n1) Droplet public IP\n2) Path to compose file\n",
	Run: func(cmd *cobra.Command, args []string) {
	    // get CLI Args
        dropletIP := args[0]
        pathToCompose := args[1]

        // run dropletactions command
        err := dropletactions.StartCompose(dropletIP, pathToCompose)
        if err != nil {
            log.Fatalln("Error running commands on Droplet: ", err)
        }

        // no error command was successful
        fmt.Println("Command successful. App now running at: ", dropletIP)
        
	    
	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(StartComposeCmd)
}
