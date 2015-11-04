// This package works with Rackspace Public Cloud Files objects
package objectstorageobjects

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/rackspace/objectstorage/v1/objects"
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
    cli.StringFlag{
      Name: "container",
      Usage: "container name to list objects",
    },
  }
}

// print list of cloud files containers to stdout
func List(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    containerName := c.Args().First()
    if c.String("container") != "" { containerName = c.String("container") }

    // step 1, set up auth options
    ao := gophercloud.AuthOptions{
      Username: user,
      APIKey: key,
    }
    // step 2, rax auth to get back provider instance
    provider, err := rackspace.AuthenticatedClient(ao)
    if err != nil { fmt.Println(err) }

    // set rax region
    serviceClient, err := rackspace.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
      Region: region,
    })
    if err != nil { fmt.Println(err) }

    // _, err := containers.Get(serviceClient, "{containerName}").ExtractMetadata()

    err3 := objects.List(serviceClient, containerName, nil).EachPage(func(page pagination.Page) (bool, error) {
      objectList, err4 := objects.ExtractNames(page)
      // https://github.com/rackspace/gophercloud/blob/master/openstack/blockstorage/v1/volumes/results.go
      for _, o := range objectList {
        fmt.Println(o)
      }
      if err4 != nil { fmt.Println(err4) }
      return true, nil
    })
    if err3 != nil { fmt.Println(err3) }
}
