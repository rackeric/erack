package objectstorageobjects

import (
  "fmt"
  "strings"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/rackspace/objectstorage/v1/objects"
)

// get list of flags for cli.go subcommand
func GetDeleteAllFlags() []cli.Flag {
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
      Name: "concurrency",
      Usage: "set TRUE to use go concurrency (really fast but misses objects)",
    },
  }
}

// delete all cloud files object from container
func DeleteAll(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    containerName := c.String("container")
    useConcurrency := c.String("concurrency")

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

    // get objects List and delete each
    err3 := objects.List(serviceClient, containerName, nil).EachPage(func(page pagination.Page) (bool, error) {
      objectList, err4 := objects.ExtractNames(page)
      // https://github.com/rackspace/gophercloud/blob/master/openstack/blockstorage/v1/volumes/results.go
      for _, objectName := range objectList {
        // now delete each object here
        // fmt.Println(objName)
        if strings.Compare(useConcurrency, "TRUE") == 0 {
          go deleteObj(serviceClient, containerName, objectName)
        } else {
          deleteObj(serviceClient, containerName, objectName)
        }
      }
      if err4 != nil { fmt.Println(err4) }
      return true, nil
    })
    if err3 != nil { fmt.Println(err3) }
}

func deleteObj(serviceClient *gophercloud.ServiceClient, containerName string, objectName string) {
  _, err2 := objects.Delete(serviceClient, containerName, objectName, nil).Extract()
  if err2 != nil { fmt.Println(err2) }
  fmt.Printf("%v Deleted.", objectName)
}
