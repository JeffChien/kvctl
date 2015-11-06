package command

import (
	"github.com/codegangsta/cli"
)

var (
	Cat   CatCommand
	Touch TouchCommand
	Rm    RmCommand
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
}
