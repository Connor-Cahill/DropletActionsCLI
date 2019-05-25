package cmd

import (
	"fmt"
	"log"
	"bytes"
	"net/http"

	"github.com/connor-cahill/dropletActionsCLI/dropletauth"
	"github.com/spf13/cobra"
)

// SignUpCmd signs new user up and creates user in Auth api
var SignUpCmd = &cobra.Command{
	Use:   "signUp",
    Short: "If users first time using CLI will sign them up and create user in authentication API, returns JWT Token.",
	Run: func(cmd *cobra.Command, args []string) {
	    // get credentials from CLI input
	    // using droplet auth package
        credentialsMap, err := dropletauth.SignUp()
        if err != nil {
            log.Fatalln("Error getting credentials from input: ", err)
        }
        // TODO: change from making request to local host to Production url 
        url := "localhost:3000/api"  // url making sign up request too
        // make request to auth api
        resp, err := http.Post(url, "application/json", credentialsMap)
        if err != nil {
            log.Fatalln("Error making request to Auth API: ", err)
        }
        defer resp.Body.Close()  // close the request body
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalln("Error reading response from auth api: ", err)
        }
        fmt.Println("RES BODY: ", string(body))
	},
}

// Mounts the command onto the root command
func init() {
	RootCmd.AddCommand(StartComposeCmd)
}
