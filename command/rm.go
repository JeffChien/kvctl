package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/jeffchien/kvctl/lib/storage"
)

type RmCommand cli.Command

func (m *RmCommand) run(c *cli.Context) {
	kv, err := storage.NewBackend(c.GlobalString("backend"))
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(c.Args()) == 0 {
		fmt.Println(fmt.Errorf("at leat one path"))
		return
	}
	for _, v := range c.Args() {
		if v[len(v)-1] == '/' {
			if !c.Bool("recursive") {
				fmt.Println("cannot remove `%s`: Is a directory", v)
				continue
			}
			err = kv.DeleteTree(v)
		} else {
			err = kv.Delete(v)
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
