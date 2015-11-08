package command

import (
	"github.com/codegangsta/cli"
)

var (
	Cat   CatCommand
	Touch TouchCommand
	Rm    RmCommand
	Mkdir MkdirCommand
)

func init() {
	Cat = CatCommand{
		Name:   "cat",
		Usage:  "print content of path.",
		Action: Cat.run,
	}
	Touch = TouchCommand{
		Name:   "touch",
		Usage:  "touch a key with value.",
		Action: Touch.run,
	}
	Rm = RmCommand{
		Name:  "rm",
		Usage: "remove keys or dirs",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "recursive,r,R",
				Usage: "remove directorieds and their content recursively.",
			},
		},
		Action: Rm.run,
	}
	Mkdir = MkdirCommand{
		Name:  "mkdir",
		Usage: "create directory",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "parent,p",
				Usage: "create parent directories if not exist.",
			},
		},
		Action: Mkdir.run,
	}
}
