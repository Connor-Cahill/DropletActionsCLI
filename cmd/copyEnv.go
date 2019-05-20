package cmd

import (
	"fmt"
	"log"

	dropletactions "github.com/connor-cahill/dropletAutomationCLI/dropletActions"
	"github.com/spf13/cobra"
)

// CopyEnvCmd secure copies over your .env file onto DO Droplet
// Takes in the following arguments IN ORDER:
// 1) Path to .env file: given as string
// 2) DO Droplet IP: as string
// 3) dirName: directory name/path on Droplet you want .env copied into
var CopyEnvCmd = &cobra.Command{
	Use:   "copyEnv",
    Short: "Secure copies .env file onto DO Droplet. Takes in the following args IN ORDER:\n 1) Path to .env: given as string\n 2) Droplet Public IP: as string\n 3) dirName: dir name/path on Droplet where you want .env copied into",
	Run: func(cmd *cobra.Command, args []string) {
	    // get CLI Args
        pathToEnv := args[0]
        dropletIP := args[1]
        dirName := args[2]

        // call droplet action
        err := dropletactions.CopyEnv(pathToEnv, dropletIP, dirName)
        if err != nil {
            log.Fatalln("Error copying over .env: ", err)
        }
	    
	    // command was successful
	    fmt.Println(".env Successfully copied onto Droplet.")

	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(CopyEnvCmd)
}
