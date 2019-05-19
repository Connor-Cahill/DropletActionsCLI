package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ListCmd is CLI Command that will list out information about
// all docker droplets
// TODO: How to handle when people have lots of droplet
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List information about all your Docker Droplets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here is the args: ", args)

	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(ListCmd)
}
