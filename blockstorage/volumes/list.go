package blockstoragevolumes

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/rackspace/blockstorage/v1/volumes"
)

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
    serviceClient, err2 := rackspace.NewBlockStorageV1(provider, gophercloud.EndpointOpts{
      Region: region,
    })
    if err2 != nil { fmt.Println(err2) }

    err3 := volumes.List(serviceClient).EachPage(func(page pagination.Page) (bool, error) {
      volumeList, err4 := volumes.ExtractVolumes(page)
      // https://github.com/rackspace/gophercloud/blob/master/openstack/blockstorage/v1/volumes/results.go
      for _, v := range volumeList {
        fmt.Println("Name: ", v.Name)
        fmt.Println("ID: ", v.ID)
        fmt.Println("Size: ", v.Size)
        fmt.Println("Status: ", v.Status)
        fmt.Println("Type: ", v.VolumeType)
        fmt.Println("Created: ", v.CreatedAt)
        fmt.Println("\n")
      }
      if err4 != nil { fmt.Println(err4) }
      return true, nil
    })
    if err3 != nil { fmt.Println(err3) }
}
