package command

import (
	"fmt"

	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/codegangsta/cli"
)

type CatCommand cli.Command

func (m *CatCommand) run(c *cli.Context) {
	kv, err := storage.New(c.GlobalString("backend"))
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(c.Args()) == 0 {
		fmt.Println(fmt.Errorf("at leat one path"))
		return
	}
	cmd, ok := kv.(lib.Command)
	if !ok {
		fmt.Println("not a command")
		return
	}
	for _, v := range c.Args() {
		pair, err := cmd.Cat(v)
		if err != nil {
			fmt.Println(lib.PrefixError(v, err))
			continue
		}
		fmt.Println(string(pair.Value))
	}
}
