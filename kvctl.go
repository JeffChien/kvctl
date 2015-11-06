package main

import (
	"github.com/codegangsta/cli"
	"github.com/jeffchien/kvctl/command"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "kvctl"
	app.Usage = "manage kv"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "backend",
			Value:  "consul://127.0.0.1:8500",
			Usage:  "backend kv storage path",
			EnvVar: "KVCTL_BACKEND",
		},
	}

	app.Commands = []cli.Command{
		cli.Command(command.Cat),
		cli.Command(command.Touch),
	}
	app.Run(os.Args)
}
