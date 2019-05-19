package dropletactions

import (
	"context"
	"fmt"
	"os/exec"

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

	return droplets, nil
}

// Create creates new digital ocean droplet must input given params:
// Name: the name of droplet
// TODO: give region options to user
// Region: region droplet will be located
// size: size / money of droplet
// image: build of droplet
// SSHKeys: what ssh keys will be configged with droplet
func Create(client *godo.Client, name string) (int, error) {

	// uses Ubuntu Droplet image
	image := godo.DropletCreateImage{
		Slug: "ubuntu-16-04-x64",
	}

	// size is to the $5/month plan by defualt currently
	// can change later but need to give users way to easily input
	size := "s-1vcpu-1gb"

	// currently have region set to SF2
	region := "sfo2"

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
		return 0, err
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
		return 0, err
	}

	return droplet.ID, nil
}

// Get returns existing droplet given droplet ID
// Must pass in authenticated client
func Get(client *godo.Client, id int) (*godo.Droplet, error) {
	// create context for digital ocean request
	ctx := context.TODO()

	// Get existing droplet using id
	droplet, _, err := client.Droplets.Get(ctx, id)
	if err != nil {
		return &godo.Droplet{}, err
	}

	return droplet, nil
}

// Delete takes in authenticated digital ocean client
// and droplet id and removes droplet
func Delete(client *godo.Client, ID int) error {
	// create context for droplet request
	ctx := context.TODO()
	// remove droplet using ID from args
	_, err := client.Droplets.Delete(ctx, ID)
	if err != nil {
		return err
	}

	// droplet successfully deleted
	// no error returned
	return nil
}

// SSHController controlls
// where user can ssh
type SSHController struct {
	User string
	IP   string
}

// DockerSetup runs script to setup docker droplet with
// docker and docker-compose CLI
func DockerSetup(ip string) error {
	cmd := exec.Command("ssh", "root@"+ip+" 'bash -s' < curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh ")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// TestExec is for testing someting
// TODO: Remove this function it is only for testing
func TestExec() error {
	filePath, err := exec.LookPath("fresh-docker-droplet.sh")
	if err != nil {
		return err
	}
	fmt.Println("PATH: ", filePath)
	return nil
}
