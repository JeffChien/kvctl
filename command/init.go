package command

import (
	"github.com/codegangsta/cli"
)

var (
	Cat   CatCommand
)

func init() {
	Cat = CatCommand{
		Name:   "cat",
		Usage:  "print content of path.",
		Action: Cat.run,
	}
}
