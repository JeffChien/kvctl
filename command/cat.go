package command

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/docker/libkv/store"
	"github.com/JeffChien/kvctl/lib/storage"
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
		pair, err := m.cat(kv, v)
		if err != nil {
			fmt.Println(PrefixError(v, err))
			continue
		}
		fmt.Println(string(pair.Value))
	}
}

func (m *CatCommand) cat(kv store.Store, path string) (*store.KVPair, error) {
	if path[len(path)-1] == '/' {
		return nil, errors.New("Is a directory")
	}
	return kv.Get(path)
}
