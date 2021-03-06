// This package works with Rackspace Public Cloud Servers flavors
package serversflavorlist

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/flavors"
)

// get list of flags for cli.go subcommand
func GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "user, u",
			Usage:  "set api username",
			EnvVar: "USERNAME",
		},
		cli.StringFlag{
			Name:   "key, k",
			Usage:  "set api key",
			EnvVar: "APIKEY",
		},
		cli.StringFlag{
			Name:   "region, r",
			Usage:  "set api region",
			EnvVar: "REGION",
		},
	}
}

// print list of cloud servers flavors to stdout
func Get(c *cli.Context) {
	// assign vars from cli args
	user := c.String("user")
	key := c.String("key")
	region := c.String("region")

	// step 1, set up auth options
	ao := gophercloud.AuthOptions{
		Username: user,
		APIKey:   key,
	}
	// step 2, rax auth to get back provider instance
	provider, err := rackspace.AuthenticatedClient(ao)
	if err != nil {	fmt.Println(err) }

	// set rax region
	serviceClient, err := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
	if err != nil {	fmt.Println(err) }

	opts := flavors.ListOpts{}
	cmdErr := flavors.ListDetail(serviceClient, opts).EachPage(func(page pagination.Page) (bool, error) {
		flavors, err := flavors.ExtractFlavors(page)
		if err != nil {	fmt.Println(err) }
		// Use the page of []flavors.Flavor
		// https://github.com/rackspace/gophercloud/blob/master/openstack/compute/v2/flavors/results.go
		for _, f := range flavors {
			fmt.Println("ID: ", f.ID)
			fmt.Println("Name: ", f.Name)
			fmt.Println("Disk: ", f.Disk)
			fmt.Println("RAM: ", f.RAM)
			fmt.Println("VCPUs: ", f.VCPUs)
			fmt.Println("\n")
		}
		return true, nil
	})
	if cmdErr != nil { fmt.Println(cmdErr) }
}
