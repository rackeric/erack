// Main package for project erack, a command
// line app for the Rackspace Public Cloud
package main

import (
	"github.com/codegangsta/cli"
	"github.com/rackeric/erack/blockstorage/volumes"
	"github.com/rackeric/erack/networks"
	"github.com/rackeric/erack/servers/flavor"
	"github.com/rackeric/erack/servers/image"
	"github.com/rackeric/erack/servers/instance"
	"os"
)

// main app function starts here
func main() {
	// set up codegangsta's cli, shit rocks
	app := cli.NewApp()
	app.Name = "erack"
	app.Usage = "CLI to the Rackspace Public Cloud"
	app.Version = "0.0.1"

	// set any -flags to the command
	// probably should remove this now?
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "user, u",
			Usage:  "set api username",
			EnvVar: "USERNAME",
		},
		cli.StringFlag{
			Name:   "key, k",
			Usage:  "set api key",
			EnvVar: "APIKEY",
		},
		cli.StringFlag{
			Name:   "region, r",
			Usage:  "set api region",
			EnvVar: "REGION",
		},
	}

	// remove this now?
	app.Flags = myFlags

	// now set subcommands
	app.Commands = []cli.Command{
		// for Cloud Servers
		{
			Name:  "servers",
			Usage: "options for Cloud Servers",
			Subcommands: []cli.Command{
				{
					Name:  "instance",
					Usage: "server instance commands",
					Subcommands: []cli.Command{
						{
							Name:  "list",
							Usage: "list server instances",
							Flags: serversinstance.GetListFlags(),
							Action: func(c *cli.Context) {
								serversinstance.GetList(c)
							},
						},
						{
							Name:  "details",
							Usage: "Details about a Cloud Server in the Rackspace Public Cloud.",
							Flags: serversinstance.GetDetailsFlags(),
							Action: func(c *cli.Context) {
								serversinstance.Details(c)
							},
						},
						{
							Name:  "create",
							Usage: "Create a Cloud Server in the Rackspace Public Cloud.",
							Flags: serversinstance.GetCreateFlags(),
							Action: func(c *cli.Context) {
								serversinstance.Create(c)
							},
						},
						{
							Name:  "delete",
							Usage: "Delete a Cloud Server in the Rackspace Public Cloud.",
							Flags: serversinstance.GetDeleteFlags(),
							Action: func(c *cli.Context) {
								serversinstance.Delete(c)
							},
						},
					},
				},
				{
					Name:  "flavor",
					Usage: "options for Cloud Servers flavor",
					Subcommands: []cli.Command{
						{
							Name:  "list",
							Usage: "Return a list of Cloud Server flavors for chosen region.",
							Flags: serversflavorlist.GetFlags(),
							Action: func(c *cli.Context) {
								serversflavorlist.Get(c)
							},
						},
					},
				},
				{
					Name:  "image",
					Usage: "options for Cloud Servers image",
					Subcommands: []cli.Command{
						{
							Name:  "list",
							Usage: "Return a list of Cloud Server images for chosen region.",
							Flags: serversimagelist.GetFlags(),
							Action: func(c *cli.Context) {
								serversimagelist.Get(c)
							},
						},
					},
				},
			},
		},
		{
			Name:  "blockstorage",
			Usage: "options for Cloud Block Storage",
			Subcommands: []cli.Command{
				{
					Name:  "volumes",
					Usage: "options for volumes",
					Subcommands: []cli.Command{
						{
							Name:  "list",
							Usage: "Return a list of Cloud Block Storage volumes.",
							Flags: blockstoragevolumes.GetListFlags(),
							Action: func(c *cli.Context) {
								blockstoragevolumes.GetList(c)
							},
						},
						{
							Name:  "show",
							Usage: "Show details on a Cloud Block Storage volume.",
							Flags: blockstoragevolumes.GetShowFlags(),
							Action: func(c *cli.Context) {
								blockstoragevolumes.Show(c)
							},
						},
						{
							Name:  "create",
							Usage: "Create a Cloud Block Storage volume.",
							Flags: blockstoragevolumes.GetCreateFlags(),
							Action: func(c *cli.Context) {
								blockstoragevolumes.Create(c)
							},
						},
						{
							Name:  "delete",
							Usage: "Delete a Cloud Block Storage volume.",
							Flags: blockstoragevolumes.GetDeleteFlags(),
							Action: func(c *cli.Context) {
								blockstoragevolumes.Delete(c)
							},
						},
						{
							Name:  "types",
							Usage: "Return a list of Cloud Block Storage volume types.",
							Flags: blockstoragevolumes.GetTypesFlags(),
							Action: func(c *cli.Context) {
								blockstoragevolumes.GetTypes(c)
							},
						},
					},
				},
			},
		},
		{
			Name:  "networks",
			Usage: "options for Cloud Networks",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "networks list",
					Flags: networks.GetListFlags(),
					Action: func(c *cli.Context) {
						networks.List(c)
					},
				},
			},
		},
	}

	// finally run the cli app
	app.Run(os.Args)
} // from func main
