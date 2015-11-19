package command

import (
	"fmt"
	"sort"

	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/codegangsta/cli"
	"github.com/docker/libkv/store"
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
	cmd, ok := kv.(lib.Command)
	if !ok {
		fmt.Println("not a command")
		return
	}
	for _, v := range paths {
		pairs, err := cmd.Ls(v)
		if err != nil && v != "" {
			fmt.Println(lib.PrefixError(v, err))
			continue
		}
		sort.Sort(byDictionary(pairs))
		for _, vv := range pairs {
			fmt.Println(vv.Key)
		}
	}
}

type byDictionary []*store.KVPair

func (m byDictionary) Len() int           { return len(m) }
func (m byDictionary) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m byDictionary) Less(i, j int) bool { return m[i].Key < m[j].Key }
