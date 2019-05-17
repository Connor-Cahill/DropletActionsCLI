package cmd

import (
    "github.com/spf13/cobra"
)


// RootCmd for our droplet automation CLI tool
var RootCmd = &cobra.Command{
    Use: "dockerdroplet",
    Short: "A CLI to help automate deploying projects on Dgital Ocean Droplets leveraging Docker.",
}
