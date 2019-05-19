package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/connor-cahill/dropletAutomationCLI/cmd"
	"github.com/digitalocean/godo"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// Init function will load in env vars from .env file
func init() {
	// Setup env vars from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error setting up env vars: ", err)
	}
	fmt.Println("Loading vars")
}

// TokenSource source for digital ocean auth token
type TokenSource struct {
	AccessToken string
}

// Token Auth setup taken from docs
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
	// note digital ocean PAT key must be in .env
	tokenSource := &TokenSource{
		AccessToken: os.Getenv("DIGITAL_OCEAN_KEY"),
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	// this is the authenticated client
	// must be passed to all other functions
	client := godo.NewClient(oauthClient)
	_ = client

	// Execute the CLI package
	// and mount the root command
	// TODO: how can I pass the client obj to the CLI?
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatalln("Error starting the CLI: ", err)
	}
}
