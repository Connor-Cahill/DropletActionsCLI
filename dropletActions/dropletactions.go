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

// Create creates new digital ocean droplet must input given params:
// Name: the name of droplet
// TODO: give region options to user
// Region: region droplet will be located
// size: size / money of droplet
// image: build of droplet
// SSHKeys: what ssh keys will be configged with droplet
func Create(client *godo.Client, name string, region string, size string, image godo.DropletCreateImage, SSHKeys []godo.DropletCreateSSHKey) (string, error) {
	// request to create new droplet
	dropletRequest := &godo.DropletCreateRequest{
		Name:    name,
		Region:  region,
		Size:    size,
		Image:   image,
		SSHKeys: SSHKeys,
		IPv6:    false,
	}

	// create context
	ctx := context.TODO()

	droplet, _, err := client.Droplets.Create(ctx, dropletRequest)
	if err != nil {
		return "", err
	}

	publicIP, err := droplet.PublicIPv4()
	if err != nil {
		return "", err
	}

	return publicIP, nil
}
