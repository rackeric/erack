package serversinstance

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
)

// get list of flags for cli.go subcommand
func GetListFlags() []cli.Flag {
  return []cli.Flag{
    cli.StringFlag{
      Name: "user, u",
      Usage: "set api username",
      EnvVar: "USERNAME",
    },
    cli.StringFlag{
      Name: "key, k",
      Usage: "set api key",
      EnvVar: "APIKEY",
    },
    cli.StringFlag{
      Name: "region, r",
      Usage: "set api region",
      EnvVar: "REGION",
    },
  }
}

func GetList(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")

    // step 1, set up auth options
    ao := gophercloud.AuthOptions{
      Username: user,
      APIKey: key,
    }
    // step 2, rax auth to get back provider instance
    provider, err := rackspace.AuthenticatedClient(ao)
    if err != nil { fmt.Println(err) }

    // set rax region
    serviceClient, err2 := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
      Region: region,
    })
    if err2 != nil { fmt.Println(err2) }

    // Retrieve a pager (i.e. a paginated collection)
    server_pager := servers.List(serviceClient, nil)

    // Define an anonymous function to be executed on each page's iteration
    err6 := server_pager.EachPage(func(page pagination.Page) (bool, error) {
      serverList, err7 := servers.ExtractServers(page)
      if err7 != nil { fmt.Println(err7) }
      for _, s := range serverList {
        // "s" will be a servers.Server
        // https://github.com/rackspace/gophercloud/blob/master/openstack/compute/v2/servers/results.go
        fmt.Println("Name: ", s.Name)
        fmt.Println("UUID: ", s.ID)
        fmt.Println("Created: ", s.Created)
        fmt.Println("Updated: ", s.Updated)
        fmt.Println("Progress: ", s.Progress)
        fmt.Println("HostID: ", s.HostID)
        fmt.Println("Image: ", s.Image["id"])
        fmt.Println("Flavor ID: ", s.Flavor["id"])
        fmt.Println("Status: ", s.Status)
        fmt.Println("Access: ", s.AccessIPv4)
        fmt.Println("Public: ", s.Addresses["public"])
        fmt.Println("Addresses: ", s.Addresses)
        fmt.Println("\n")
      }
      return true, nil
    })
    if err6 != nil { fmt.Println(err6) }
}
