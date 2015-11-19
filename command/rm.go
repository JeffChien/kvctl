package command

import (
	"fmt"

	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/codegangsta/cli"
)

type RmCommand cli.Command

func (m *RmCommand) run(c *cli.Context) {
	kv, err := storage.New(c.GlobalString("backend"))
	if err != nil {
		fmt.Println(err)
		return
	}
	paths := []string{}
	if len(c.Args()) == 0 {
		paths = append(paths, "")
	} else {
		paths = append(paths, c.Args()...)
	}
	cmd, ok := kv.(lib.Command)
	if !ok {
		fmt.Println("not a command")
		return
	}
	for _, v := range paths {
		err = cmd.Rm(v, c.Bool("recursive"))
		if err != nil {
			fmt.Println(lib.PrefixError(v, err))
		}
	}
}
