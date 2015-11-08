package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/docker/libkv/store"
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
		err = m.rm(kv, v, c.Bool("recursive"))
		if err != nil {
			fmt.Println(PrefixError(v, err))
		}
	}
}

func (m *RmCommand) rm(kv store.Store, path string, recursive bool) error {
	var err error
	if path[len(path)-1] == '/' {
		if !recursive {
			return fmt.Errorf("Is a directory")
		}
		err = kv.DeleteTree(path)
	} else {
		err = kv.Delete(path)
	}
	return err
}
