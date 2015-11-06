package command

import (
	"github.com/codegangsta/cli"
)

var (
	Cat   CatCommand
	Touch TouchCommand
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
}
