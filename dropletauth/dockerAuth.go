package dropletauth

import (
	"context"
	"os"
	"errors"
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


// SignUp gets users credentials for signing up
// returns Map or string string key, value pair
// to send with request body to Auth API
func SignUp() (map[string]string, error) {
    // create new reader to get Stdin
    reader := bufio.NewReader(os.Stdin)
    // ask user for email to sign up
    fmt.Println("Please Enter An Email: ")
    email, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }
    // ask user for a password
    fmt.Println("Please Enter A Password: ")
    password, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }
    // ask user for confirmation password
    fmt.Println("Confirm Your Password: ")
    confirmPassword, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }
    
    // make sure passwords matchup
    if password != confirmPassword {
        return nil, errors.New("the passwords you entered did not match up")
    }

    // build return map
    credentialsMap := map[string]string{
        "email": email,
        "password": password,
    }

    // return the credentials map
    return credentialsMap, nil
}

// SignIn gets users email and password from CLI 
// and returns in map to be sent to node auth api
// and authenticate user
func SignIn() (map[string]string, error) {
    // create reader to parse CLI parameters
    reader := bufio.NewReader(os.Stdin)
    // ask user for email
    fmt.Println("Please enter your email: ")
    email, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }
    // ask user for password
    fmt.Println("Please enter your password: ")
    password, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }
    // buildout the credential map
    credentialsMap := map[string]string{
        "email": email,
        "password": password,
    }
    // return credentials map
    return credentialsMap, nil
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
