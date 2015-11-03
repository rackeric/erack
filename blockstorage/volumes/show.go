package blockstoragevolumes

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  //"github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/rackspace/blockstorage/v1/volumes"
)

// get list of flags for cli.go subcommand
func GetShowFlags() []cli.Flag {
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
      Usage: "set uuid of volume to get details",
    },
  }
}

// print list of cloud block storage volumes to stdout
func Show(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    volumeId := c.Args().First()
    if c.String("uuid") != "" { volumeId = c.String("uuid") }

    // step 1, set up auth options
    ao := gophercloud.AuthOptions{
      Username: user,
      APIKey: key,
    }
    // step 2, rax auth to get back provider instance
    provider, err := rackspace.AuthenticatedClient(ao)
    if err != nil { fmt.Println(err) }

    // set rax region
    serviceClient, err2 := rackspace.NewBlockStorageV1(provider, gophercloud.EndpointOpts{
      Region: region,
    })
    if err2 != nil { fmt.Println(err2) }

    v, err3 := volumes.Get(serviceClient, volumeId).Extract()
    if err3 != nil { fmt.Println(err3) }
    fmt.Println("Name: ", v.Name)
    fmt.Println("ID: ", v.ID)
    fmt.Println("Size: ", v.Size)
    fmt.Println("Status: ", v.Status)
    fmt.Println("Type: ", v.VolumeType)
    fmt.Println("Created: ", v.CreatedAt)
}
