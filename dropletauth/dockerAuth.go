package dropletauth

import (
	"context"
	"os"
	"fmt"
	"bufio"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// TokenSource struct holds the
// Digital Ocean Personal Access Token
type TokenSource struct {
	AccessToken string
}


// GetToken will get token from Command line input and set env
func GetToken() error { 
    reader := bufio.NewReader(os.Stdin)
    // Ask user for Digital Ocean Key
    fmt.Println("Input Digital Ocean API Key: ")
    key, err := reader.ReadString('\n')
    if err != nil {
        return err
    }
    // sets inputted digital ocean API Key to env vars
    os.Setenv("DIGITAL_OCEAN_KEY", key)
    return nil
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

// Auth uses DO Personal Access Token and returns
// authenticated client to then use with
// Droplet automation commands
func Auth() *godo.Client {
	tokenSource := &TokenSource{
		AccessToken: os.Getenv("DIGITAL_OCEAN_KEY"),
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	// returing authenticated client object
	client := godo.NewClient(oauthClient)

	return client

}
