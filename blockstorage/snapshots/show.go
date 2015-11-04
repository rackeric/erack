package blockstoragesnapshots

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/rackspace/blockstorage/v1/snapshots"
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
      Usage: "set uuid of snapshot to show details",
    },
  }
}

// print details of a cloud block storage snapshots to stdout
func Show(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    snapshotId := c.Args().First()
    if c.String("uuid") != "" { snapshotId = c.String("uuid") }

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

    s, err3 := snapshots.Get(serviceClient, snapshotId).Extract()
    if err3 != nil { fmt.Println(err3) }

    fmt.Println("Name: ", s.Name)
    fmt.Println("ID: ", s.ID)
    fmt.Println("Size: ", s.Size)
    fmt.Println("Status: ", s.Status)
    fmt.Println("Progress: ", s.Progress)
    fmt.Println("Parent: ", s.VolumeID)
    fmt.Println("Created: ", s.CreatedAt)
    fmt.Println("\n")
}
