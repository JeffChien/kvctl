package command

import (
	"fmt"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/codegangsta/cli"
	"github.com/docker/libkv/store"
	"sort"
)

type LsCommand cli.Command

func (m *LsCommand) run(c *cli.Context) {
	var paths []string = []string(c.Args())
	kv, err := storage.New(c.GlobalString("backend"))
	if err != nil {
		fmt.Println(err)
		return
	}
	if !c.Args().Present() {
		paths = append(paths, "")
	}
	for _, v := range paths {
		pairs, err := m.ls(kv, v)
		if err != nil && v != "" {
			fmt.Println(PrefixError(v, err))
			continue
		}
		for _, vv := range pairs {
			fmt.Println(vv.Key)
		}
	}
}

func (m *LsCommand) ls(kv store.Store, path string) ([]*store.KVPair, error) {
	var err error
	pairs, err := kv.List(path)
	if err != nil {
		return nil, err
	}
	sort.Sort(byDictionary(pairs))
	return pairs, nil
}

type byDictionary []*store.KVPair

func (m byDictionary) Len() int           { return len(m) }
func (m byDictionary) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m byDictionary) Less(i, j int) bool { return m[i].Key < m[j].Key }
