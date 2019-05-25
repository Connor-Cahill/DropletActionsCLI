package cmd

import (
	"fmt"
	"log"

	dropletactions "github.com/connor-cahill/dropletActionsCLI/dropletActions"
	"github.com/connor-cahill/dropletAutomationCLI/dropletauth"
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
		client := dropletauth.Auth()
		droplets, err := dropletactions.Index(client)
		if err != nil {
			log.Fatalln("Error listing all Droplet information: ", err)
		}

		for _, droplet := range droplets {
			fmt.Println("_____________________________________________")
			fmt.Println("Droplet: ", droplet.Name)
			fmt.Println("Droplet ID: ", droplet.ID)
			publicIP, err := droplet.PublicIPv4()
			if err != nil {
				log.Fatalln("Error getting droplet information: ", err)
			}
			fmt.Println("Droplet IP: ", publicIP)
		}
		fmt.Println("_____________________________________________")
	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(ListCmd)
}
