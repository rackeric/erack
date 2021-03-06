# erack
[![Build Status](http://104.130.11.192:8080/buildStatus/icon?job=test1)](http://104.130.11.192:8080/job/test1/)
[![GoDoc](https://godoc.org/github.com/rackeric/erack?status.svg)](https://godoc.org/github.com/rackeric/erack)

Another CLI to the Rackspace Public Cloud written in Go and using a couple of really cool libraries.

- [The Go Programming Language](https://golang.org/)
- [gophercloud](http://gophercloud.io/)
- [cli.go](https://github.com/codegangsta/cli)

## Requirements
```
$ go get github.com/codegangsta/cli
$ go get github.com/rackspace/gophercloud
```

## Authentication
You can authenticate by setting your username, API key and region using the command line options.
```
$ erack servers instance list --user <username> --key <api_key> --region <region>
```

## Environment Variables
You can also set your `USERNAME`, `APIKEY` and `REGION` by using environment variables instead of using command line options.
```
$ export USERNAME=<username>
$ export APIKEY=<api_key>
$ export REGION=<region>
```

## Built in help
You can see help info for each subcommand by appending help, --help or -h to a subcommand.
```
$ erack servers instance help
NAME:
   erack servers instance - server instance commands

USAGE:
   erack servers instance command [command options] [arguments...]

COMMANDS:
   list         list server instances
   details      Details about a Cloud Server in the Rackspace Public Cloud.
   create       Create a Cloud Server in the Rackspace Public Cloud.
   delete       Delete a Cloud Server in the Rackspace Public Cloud.
   help, h      Shows a list of commands or help for one command

OPTIONS:
   --help, -h   show help
```
