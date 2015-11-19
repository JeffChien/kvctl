package command

import (
	"fmt"
	"time"

	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/codegangsta/cli"
)

type MkdirCommand cli.Command

type mkdirOption struct {
	Parent bool
	TTL    time.Duration
}

// only available in etcd
func (m *MkdirCommand) run(c *cli.Context) {
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
		err := cmd.Mkdir(v, &lib.MkdirOption{
			Parent: c.Bool("parent"),
		})
		if err != nil {
			fmt.Println(lib.PrefixError(v, err))
			continue
		}
	}
}
