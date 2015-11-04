// This package works with Rackspace Public Cloud Block Storage snapshots
package blockstoragesnapshots

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/rackspace/blockstorage/v1/snapshots"
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

// print list of cloud block storage snapshots to stdout
func List(c *cli.Context) {
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
    serviceClient, err2 := rackspace.NewBlockStorageV1(provider, gophercloud.EndpointOpts{
      Region: region,
    })
    if err2 != nil { fmt.Println(err2) }

    err3 := snapshots.List(serviceClient).EachPage(func(page pagination.Page) (bool, error) {
      snapshotList, err4 := snapshots.ExtractSnapshots(page)
      for _, s := range snapshotList {
        fmt.Println("Name: ", s.Name)
        fmt.Println("ID: ", s.ID)
        fmt.Println("Size: ", s.Size)
        fmt.Println("Status: ", s.Status)
        //fmt.Println("Type: ", s.VolumeType)  // where is this?
        fmt.Println("Parent: ", s.VolumeID)
        fmt.Println("Created: ", s.CreatedAt)
        fmt.Println("\n")
      }
      if err4 != nil { fmt.Println(err4) }
      return true, nil
    })
    if err3 != nil { fmt.Println(err3) }
}
