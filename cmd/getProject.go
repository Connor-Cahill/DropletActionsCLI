package cmd

import (
	"fmt"
	"log"

	dropletactions "github.com/connor-cahill/dropletAutomationCLI/dropletActions"
	"github.com/spf13/cobra"
)

// GetProjectCmd is a CLI command that will clone project repo onto digital ocean droplet
// command must be passed the following ards in order:
// 1) DO Droplet public IP 2) github clone link to repo 3) directory name you want the project cloned into on droplet
var GetProjectCmd = &cobra.Command{
	Use:   "getProject",
	Short: "Clones project onto DO Droplet must pass the following args in order: 1) DO Droplet Public IP 2) Github clone link for repo 3) directory name you want project cloned into on Droplet.",
	Run: func(cmd *cobra.Command, args []string) {
		// get authenticated DO Client

		// get droplet public ip address from args
		dropletIP := args[0]
		projectRepo := args[1]
		dirName := args[2]

		err := dropletactions.GetProject(dropletIP, projectRepo, dirName)
		if err != nil {
			log.Fatalln("Error setting up Docker on Droplet: ", err)
		}

		fmt.Println("Docker & docker-compose setup on Droplet.")
	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(GetProjectCmd)
}
