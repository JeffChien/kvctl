package command

import (
	"fmt"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/codegangsta/cli"
	"github.com/docker/libkv/store"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"path/filepath"
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

	for _, v := range c.Args() {
		err := m.touch(kv, v, data, nil)
		if err != nil {
			fmt.Println(PrefixError(v, err))
			continue
		}
	}
}

func (m *TouchCommand) touch(kv store.Store, path string, data []byte, opts *store.WriteOptions) error {
	if path[len(path)-1] == '/' {
		return fmt.Errorf("got a directory")
	}
	normalizePath := filepath.Clean(path)
	parentDir := filepath.Dir(normalizePath)
	if !(parentDir == "." || parentDir == "/") {
		//check parent dir existk
		exists, err := kv.Exists(fmt.Sprintf("%s/", parentDir))
		if err != nil {
			return err
		}
		if !exists {
			return store.ErrKeyNotFound
		}
	}
	return kv.Put(normalizePath, data, opts)
}
