package main

import (
	"os"

	"github.com/JeffChien/kvctl/command"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "kvctl"
	app.Version = kvctl_version
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
		cli.Command(command.Rm),
		cli.Command(command.Mkdir),
		cli.Command(command.Ls),
		cli.Command(command.Tar),
	}
	app.Run(os.Args)
}
