package objectstorageobjects

import (
  "fmt"
  "io/ioutil"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/rackspace/objectstorage/v1/objects"
)

// get list of flags for cli.go subcommand
func GetDownloadFlags() []cli.Flag {
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
      Usage: "container name to upload file",
    },
    cli.StringFlag{
      Name: "name",
      Usage: "name of object to download",
    },
    cli.StringFlag{
      Name: "path",
      Usage: "set local path for download",
    },
  }
}

// download file from Cloud File container
func Download(c *cli.Context) {
    // assign vars from cli args
    user := c.String("user")
    key := c.String("key")
    region := c.String("region")
    containerName := c.String("container")
    objectName := c.String("name")
    path := c.String("path")

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

    result := objects.Download(serviceClient, containerName, objectName, nil)
    content, err := result.ExtractContent()

    err3 := ioutil.WriteFile(path + "/" + objectName, []byte(content), 0644)
    if err3 != nil { fmt.Println(err3) }
    fmt.Printf("%v Downloaded.", objectName)
}
