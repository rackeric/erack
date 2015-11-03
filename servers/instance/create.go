package serversinstance

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
)

// get list of flags for cli.go subcommand
func GetCreateFlags() []cli.Flag {
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
      Name: "name, n",
      Usage: "name of new server to create",
    },
    cli.StringFlag{
      Name: "image, i",
      Usage: "image for new server",
    },
    cli.StringFlag{
      Name: "flavor, f",
      Usage: "flavor for new server",
    },
  }
}

// create cloud servers instance and print out results to stdout (eg. password)
func Create(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    // server create specific options
    newservername := c.String("name")
    imageid := c.String("image")
    flavorid := c.String("flavor")

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

    // create the cloud server
    server, err3 := servers.Create(serviceClient, servers.CreateOpts{
      Name:      newservername,
      ImageRef:  imageid,
      FlavorRef: flavorid,
    }).Extract()

    fmt.Println(server)
    if err3 != nil { fmt.Println(err3)}
}
