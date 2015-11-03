package blockstoragevolumes

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  //"github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/rackspace/blockstorage/v1/volumes"
  osvolumes "github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumes"
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
      Usage: "set name of new cloud block storage volume",
    },
    cli.StringFlag{
      Name: "size",
      Usage: "set size of new cloud block storage volume",
    },
    cli.StringFlag{
      Name: "type",
      Usage: "set type of new cloud block storage volume",
    },
  }
}

// create cloud block storage volume in rackspace public cloud
func Create(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    name := c.String("name")
    size := c.Int("size")
    var volType string
    if c.String("type") != "" {
      volType = c.String("type")
    }

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


    // https://github.com/rackspace/gophercloud/blob/master/openstack/blockstorage/v1/volumes/requests.go
    opts := osvolumes.CreateOpts{
      Name: name,
      Size: size,
      VolumeType: volType,
    }
    if volType == "" {
      opts = osvolumes.CreateOpts{
        Name: name,
        Size: size,
      }
    }

    vol, err3 := volumes.Create(serviceClient, opts).Extract()

    if err3 != nil { fmt.Println(err3) }
    fmt.Println(vol)
}
