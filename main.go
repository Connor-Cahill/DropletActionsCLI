package main

import (

    "log"
    "github.com/connor-cahill/dropletAutomationCLI/cmd"
)


func main() {

    // Execute the CLI package
    // and mount the root command
    err := cmd.RootCmd.Execute()
    if err != nil {
        log.Fatalln("Error starting the CLI: ", err)
    }
}
