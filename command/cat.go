package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/jeffchien/kvctl/lib/storage"
)

type CatCommand cli.Command

func (m *CatCommand) run(c *cli.Context) {
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
			fmt.Printf("%s: Is a directory\n", v)
			continue
		}
		pair, err := kv.Get(v)
		if err != nil {
			fmt.Println(PrefixError(v, err))
			continue
		}
		fmt.Println(string(pair.Value))
	}
}
