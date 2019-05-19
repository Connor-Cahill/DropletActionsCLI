package cmd

import (
	dropletactions "github.com/connor-cahill/dropletAutomationCLI/dropletActions"
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
		dropletactions.DockerSetup(dropletIP)
	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(SetupDockerCmd)
}