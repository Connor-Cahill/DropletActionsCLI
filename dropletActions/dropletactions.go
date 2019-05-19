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
func Create(client *godo.Client, name string) (string, error) {

	// uses Ubuntu Droplet image
	image := godo.DropletCreateImage{
		Slug: "ubuntu-16-04-x64",
	}

	// size is to the $5/month plan by defualt currently
	// can change later but need to give users way to easily input
	size := "s-1vcpu-1gb"

	// currently have region set to SF2
	region := "sf2"

	// NOTE: it currently connects all SSH keys on your account to the droplet
	// I will fix this later but for right now this is how I use DO anyway
	// options for ssh request
	sshOpt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}
	ctx := context.TODO()
	keys, _, err := client.Keys.List(ctx, sshOpt)
	if err != nil {
		return "", err
	}

	// list of SSH IDs to add to new droplet
	SSHKeys := []godo.DropletCreateSSHKey{}
	for _, key := range keys {
		sshKey := godo.DropletCreateSSHKey{
			ID: key.ID,
		}
		SSHKeys = append(SSHKeys, sshKey)
	}

	// request to create new droplet
	dropletRequest := &godo.DropletCreateRequest{
		Name:    name,
		Region:  region,
		Size:    size,
		Image:   image,
		SSHKeys: SSHKeys,
		IPv6:    false,
		Backups: false,
	}

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
