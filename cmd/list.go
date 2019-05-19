package cmd

import (
	"fmt"
	"log"

	"github.com/connor-cahill/dropletAutomationCLI/dockerauth"
	dropletactions "github.com/connor-cahill/dropletAutomationCLI/dropletActions"
	"github.com/spf13/cobra"
)

// ListCmd is CLI Command that will list out information about
// all docker droplets
// TODO: How to handle when people have lots of droplet
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List information about all your Docker Droplets",
	Run: func(cmd *cobra.Command, args []string) {
		// get authenticated client to make droplet request
		client := dockerauth.Auth()
		droplets, err := dropletactions.Index(client)
		if err != nil {
			log.Fatalln("Error listing all Droplet information: ", err)
		}

		fmt.Println("Here are your droplets: ", droplets)

	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(ListCmd)
}
