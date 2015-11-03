package networks

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/rackspace/gophercloud"
	//"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/pagination"
	//"github.com/rackspace/gophercloud/rackspace/compute/v2/networks"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
)

func GetListFlags() []cli.Flag {
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

func List(c *cli.Context) {
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
	//provider, err := rackspace.AuthenticatedClient(ao)
	provider, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		fmt.Println(err)
	}

	// set rax region
	serviceClient, err2 := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Name:   "neutron",
		Region: region,
	})
	if err2 != nil {
		fmt.Println(err2)
	}

	// Retrieve a pager (i.e. a paginated collection)
	networks_pager := networks.List(serviceClient, nil)

	// Define an anonymous function to be executed on each page's iteration
	err6 := networks_pager.EachPage(func(page pagination.Page) (bool, error) {
		networkList, err7 := networks.ExtractNetworks(page)
		if err7 != nil {
			fmt.Println(err7)
		}
		for _, n := range networkList {
			// "s" will be a servers.Server
			// https://github.com/rackspace/gophercloud/blob/master/openstack/compute/v2/servers/results.go
			fmt.Println(n)
		}
		return true, nil
	})
	if err6 != nil {
		fmt.Println(err6)
	}
}
