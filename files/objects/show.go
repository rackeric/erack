package objectstorageobjects

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/rackspace/objectstorage/v1/objects"
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
      Name: "container",
      Usage: "container name of object",
    },
    cli.StringFlag{
      Name: "name",
      Usage: "object name to show details",
    },
  }
}

// print details of cloud files object to stdout
func Show(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    containerName := c.String("container")
    objectName := c.String("name")

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

    myObject, err3 := objects.Get(serviceClient, containerName, objectName, nil).Extract()
    if err3 != nil { fmt.Println(err3) }

    fmt.Println("Name: ", objectName)
    fmt.Println("Bytes: ", myObject.ContentLength)
    fmt.Println("Type: ", myObject.ContentType)
    fmt.Println("Last Modified: ", myObject.LastModified)
}
