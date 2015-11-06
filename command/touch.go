package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/jeffchien/kvctl/lib/storage"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
)

type TouchCommand cli.Command

func (m *TouchCommand) run(c *cli.Context) {
	var data []byte
	kv, err := storage.NewBackend(c.GlobalString("backend"))
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(c.Args()) == 0 {
		fmt.Println(fmt.Errorf("at leat one path"))
		return
	}

	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		// nothing on stdin
		data = []byte("")
	} else {
		// something on stdin
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for _, v := range c.Args() {
		err := kv.Put(v, data, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("try to access path %s", v))
			continue
		}
	}
}
