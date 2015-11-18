package command

import (
	"fmt"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/codegangsta/cli"
	"github.com/docker/libkv/store"
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
	for _, v := range paths {
		err = m.rm(kv, v, c.Bool("recursive"))
		if err != nil {
			fmt.Println(PrefixError(v, err))
		}
	}
}

func (m *RmCommand) rm(kv store.Store, path string, recursive bool) error {
	var err error
	if path == "" || path[len(path)-1] == '/' {
		if !recursive {
			return fmt.Errorf("Is a directory")
		}
		err = kv.DeleteTree(path)
	} else {
		err = kv.Delete(path)
	}
	return err
}
