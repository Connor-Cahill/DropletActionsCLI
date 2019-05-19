package dropletactions

import (
	"context"

	"github.com/digitalocean/godo"
)

// Index returns list of information on all your droplets
func Index(client *godo.Client) ([]godo.Droplet, error) {
	// options for droplets request
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 20,
	}
	// create context
	ctx := context.TODO()
	// get droplet list
	droplets, _, err := client.Droplets.List(ctx, opt)
	if err != nil {
		return nil, err
	}

	// TODO: some sort of string formatting on this

	return droplets, nil
}
