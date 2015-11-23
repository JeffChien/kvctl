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

type TarCommand cli.Command

//TODO: simplify
func (m *TarCommand) run(c *cli.Context) {
	var data []byte
	var err error
	var fname string
	kv, err := storage.New(c.GlobalString("backend"))
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd, ok := kv.(lib.Command)
	if !ok {
		fmt.Println("not a command")
		return
	}
	if fname = c.String("file"); fname == "" {
		fmt.Println(fmt.Errorf("need a file"))
		return
	}
	//mutual flag
	if c.Bool("create") != c.Bool("extract") {
		if c.Bool("create") {
			if len(c.Args()) != 1 {
				fmt.Println(fmt.Errorf("need a path"))
				return
			}
			tarData, err := cmd.Dump(c.Args()[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			var f *os.File
			if fname == "-" {
				f = os.Stdout
			} else {
				if f, err = os.Create(fname); err != nil {
					fmt.Println(err)
					return
				}
				defer f.Close()
			}
			if _, err = f.Write(tarData); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			if fname == "-" {
				if !terminal.IsTerminal(int(os.Stdin.Fd())) {
					if data, err = ioutil.ReadAll(os.Stdin); err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println(data)
				}
			} else {
				if data, err = ioutil.ReadFile(fname); err != nil {
					fmt.Println(err)
					return
				}
			}
			if err = cmd.Restore(data); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
