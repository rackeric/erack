package blockstoragesnapshots

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/rackspace/blockstorage/v1/snapshots"
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
      Name: "name",
      Usage: "set name of snapshot to create",
    },
    cli.StringFlag{
      Name: "uuid",
      Usage: "the uuid of volume to take a snapshot",
    },
  }
}

// create a cloud block storage snapshots from volume uuid
func Create(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    name := c.String("name")
    volumeId := c.String("uuid")

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

    // create snapshot from volume
    opts := snapshots.CreateOpts{VolumeID: volumeId, Name: name}
    s, err3 := snapshots.Create(serviceClient, opts).Extract()

    if err3 != nil { fmt.Println(err3) }

    fmt.Println(s)
}
