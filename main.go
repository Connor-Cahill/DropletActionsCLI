package main

import (
    "os"
    "context"
    "fmt"
    "log"
    "github.com/connor-cahill/dropletAutomationCLI/cmd"
    "github.com/joho/godotenv"
    "github.com/digitalocean/godo"
    "golang.org/x/oauth2"
)



// Init function will load in env vars from .env file
func init() {
    // Setup env vars from .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalln("Error setting up env vars: ", err)
    }
}

// Digitial Ocean personal access token
// Must put in .env to authenticate
var doToken = os.Getenv("DIGITAL_OCEAN_KEY")

// TokenSource source for digital ocean auth token
type TokenSource struct {
    AccessToken string
}

// Auth setup taken from docs
// https://github.com/digitalocean/godo
// auths user using digital ocean access token
func (t *TokenSource) Token() (*oauth2.Token, error) {
    token := &oauth2.Token{
        AccessToken: t.AccessToken,
    }
    return token, nil
}

func main() {
    // TokenSource for authenticating DO user
    tokenSource := &TokenSource{
        AccessToken: doToken,
    }

    fmt.Println("DOToken: ", doToken)

    oauthClient := oauth2.NewClient(context.Background(), tokenSource)
    // this is the authenticated client
    // must be passed to all other functions
    client := godo.NewClient(oauthClient)
    _ = client
    // testing do
    // request options for getting Droplet list
    opt := &godo.ListOptions{
        Page: 1,
        PerPage: 20,
    }
    
    ctx := context.TODO()
    droplets, _, err := client.Droplets.List(ctx, opt)
    if err != nil {
        log.Fatalln("Error retrieiving droplet list: ", err)
    }
    fmt.Println("DROPLETS: ", droplets)

    // Execute the CLI package
    // and mount the root command
    // TODO: how can I pass the client obj to the CLI?
    err = cmd.RootCmd.Execute()
    if err != nil {
        log.Fatalln("Error starting the CLI: ", err)
    }
}
