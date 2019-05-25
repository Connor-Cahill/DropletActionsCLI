package cmd

import (
	"fmt"
	"log"

	dropletactions "github.com/connor-cahill/dropletActionsCLI/dropletActions"
	"github.com/spf13/cobra"
)

// SetupDockerCmd is CLI command that automates the installation
// of docker and docker-compose in your Digital Ocean Droplet
var SetupDockerCmd = &cobra.Command{
	Use:   "setupDocker",
	Short: "Given public IP address of Droplet will setup docker and docker-compose on VPS.",
	Run: func(cmd *cobra.Command, args []string) {
		// get authenticated DO Client

		// get droplet public ip address from args
		dropletIP := args[0]

		err := dropletactions.DockerSetup(dropletIP)
		if err != nil {
			log.Fatalln("Error setting up Docker on Droplet: ", err)
		}

		fmt.Println("Docker & docker-compose setup on Droplet.")
	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(SetupDockerCmd)
}
