package cmd

import (
	"log"
	"strconv"

	"github.com/connor-cahill/dropletAutomationCLI/dockerauth"
	dropletactions "github.com/connor-cahill/dropletAutomationCLI/dropletActions"
	"github.com/spf13/cobra"
)

// DeleteCmd is CLI command that will delete a Droplet running
// on your Digital Ocean account
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Destroys Digital Ocean Droplet removing it from your accont.",
	Run: func(cmd *cobra.Command, args []string) {
		// Get authenticated client to make request
		client := dockerauth.Auth()
		// Droplet id passed in with args
		dropletID := args[0]
		// convert into int id
		intID, err := strconv.Atoi(dropletID)
		if err != nil {
			log.Fatalln("Error converting ID argument to Int: ", err)
		}
		// deletes the droplet
		err = dropletactions.Delete(client, intID)
		if err != nil {
			log.Fatalln("Error destroying your Digital Ocean Droplet: ", err)
		}
	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(DeleteCmd)
}
