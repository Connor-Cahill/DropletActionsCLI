package cmd

import (
	"fmt"
	"log"

	"github.com/connor-cahill/dropletAutomationCLI/dockerauth"
	dropletactions "github.com/connor-cahill/dropletAutomationCLI/dropletActions"
	"github.com/spf13/cobra"
)

// CreateCmd is CLI Command that will list out information about
// all docker droplets
// TODO: How to handle when people have lots of droplet
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates new Droplet with base Ubuntu image, set to the $5/month plan by default. PARAM: Must enter droplet name as param.",
	Run: func(cmd *cobra.Command, args []string) {
		// get authenticated client to make droplet request
		client := dockerauth.Auth()
		dropletName := args[0]
		newDroplet, err := dropletactions.Create(client, dropletName)
		if err != nil {
			log.Fatalln("Error creating new Droplet: ", err)
		}
		fmt.Println("New Droplet IP: ", newDroplet)
	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(CreateCmd)
}
