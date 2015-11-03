package serversimagelist

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/rackspace/compute/v2/images"
)

// get list of flags for cli.go subcommand
func GetFlags() []cli.Flag {
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

// print list of cloud servers images to stdout
func Get(c *cli.Context) {
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
  serviceClient, err2 := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
    Region: region,
  })
  if err2 != nil { fmt.Println(err2) }

  //opts3 := images.ListOpts{}
  err4 := images.ListDetail(serviceClient, nil).EachPage(func (page pagination.Page) (bool, error) {
    images, err5 := images.ExtractImages(page)
    if err5 != nil { fmt.Println(err5) }
    // Use the page of []images.Image
    // https://github.com/rackspace/gophercloud/blob/master/openstack/compute/v2/images/results.go
    for _, i := range images {
      fmt.Println("Name: ", i.Name)
      fmt.Println("ID: ", i.ID)
      fmt.Println("Created: ", i.Created)
      fmt.Println("Updated: ", i.Updated)
      fmt.Println("MinDisk: ", i.MinDisk)
      fmt.Println("MinRAM: ", i.MinRAM)
      fmt.Println("Progress: ", i.Progress)
      fmt.Println("Status: ", i.Status)
      fmt.Println("\n")
    }
    return true, nil
  })
  if err4 != nil { fmt.Println(err4) }
}
