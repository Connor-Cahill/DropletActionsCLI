package cmd

import (
	"fmt"
	"log"

	dropletactions "github.com/connor-cahill/dropletActionsCLI/dropletActions"
	"github.com/spf13/cobra"
)

// DownComposeCmd will ssh into DO Droplet and run docker-compose down on compose file
// to stop docker-compose instance on droplet
// takes in DO Droplet public IP and path to compose file as arguments
var DownComposeCmd = &cobra.Command{
	Use:   "downCompose",
    Short: "Will ssh into DO Droplet and docker-compose down stopping docker-compose instance. Must pass following args IN ORDER:\n1) Droplet public IP\n2) Path to compose file\n",
	Run: func(cmd *cobra.Command, args []string) {
	    // get CLI Args
        dropletIP := args[0]
        pathToCompose := args[1]

        // run dropletactions command
        err := dropletactions.DownCompose(dropletIP, pathToCompose)
        if err != nil {
            log.Fatalln("Error running commands on Droplet: ", err)
        }

        // no error command was successful
        fmt.Println("Command successful. Docker compose instance was stopped.")
        
	    
	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(DownComposeCmd)
}
