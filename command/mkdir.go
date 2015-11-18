package command

import (
	"fmt"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/JeffChien/kvctl/util"
	"github.com/codegangsta/cli"
	"github.com/docker/libkv/store"
	"path/filepath"
	"time"
)

type MkdirCommand cli.Command
type mkdirOption struct {
	Parent bool
	TTL    time.Duration
}

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
	for _, v := range c.Args() {
		err := m.mkdir(kv, v, &mkdirOption{
			Parent: c.Bool("parent"),
		})
		if err != nil {
			fmt.Println(PrefixError(v, err))
			continue
		}
	}
}

func (m *MkdirCommand) mkdir(kv store.Store, path string, opt *mkdirOption) error {
	var err error
	var ttl time.Duration
	var createParent func(string) error
	var normalizePath string = filepath.Clean(path)
	if opt != nil {
		ttl = opt.TTL
		if opt.Parent {
			createParent = func(p string) error {
				if p == "." {
					return nil
				}
				if err = createParent(filepath.Dir(p)); err != nil {
					return err
				}
				return kv.Put(util.NormalizeDir(p), nil, &store.WriteOptions{IsDir: true})
			}
			if err = createParent(filepath.Dir(normalizePath)); err != nil {
				return err
			}
		}
	}
	parentDir := util.NormalizeDir(filepath.Dir(normalizePath))
	if parentDir != "./" {
		if parentExit, _ := kv.Exists(parentDir); !parentExit {
			return store.ErrKeyNotFound
		}
	}
	return kv.Put(util.NormalizeDir(normalizePath), nil, &store.WriteOptions{
		IsDir: true,
		TTL:   ttl,
	})
}
