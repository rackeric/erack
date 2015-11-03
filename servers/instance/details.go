package serversinstance

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
)

// get list of flags for cli.go subcommand
func GetDetailsFlags() []cli.Flag {
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
    cli.StringFlag{
      Name: "uuid",
      Usage: "Cloud Server uuid",
    },
  }
}

// print details of cloud servers instance to stdout
func Details(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    // server create specific options
    // check both options and arguments for server uuid
    serverid := c.Args().First()
    if c.String("uuid") != "" { serverid = c.String("uuid") }

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

    // get cloud server
    s, id := servers.Get(serviceClient, serverid).Extract()
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
    if id != nil { fmt.Println(id) }
}
