package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/codegangsta/cli"
	"golang.org/x/crypto/ssh/terminal"
)

type TouchCommand cli.Command

func (m *TouchCommand) run(c *cli.Context) {
	var data []byte
	kv, err := storage.New(c.GlobalString("backend"))
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

	cmd, ok := kv.(lib.Command)
	if !ok {
		fmt.Println("not a command")
		return
	}

	for _, v := range c.Args() {
		err := cmd.Touch(v, data, nil)
		if err != nil {
			fmt.Println(lib.PrefixError(v, err))
			continue
		}
	}
}
