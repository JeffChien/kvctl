package etcd

import (
	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage/general"
	"github.com/JeffChien/kvctl/util"
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/etcd"
	"path/filepath"
	"strings"
	"time"
)

type EtcdStorage struct {
	*general.GeneralStorage
}

func New(hosts []string, config *store.Config) (store.Store, error) {
	etcd.Register()

	s, err := libkv.NewStore(
		store.ETCD,
		hosts,
		config,
	)
	if err != nil {
		return nil, err
	}
	return &EtcdStorage{
		GeneralStorage: general.New(s),
	}, nil
}

func (m *EtcdStorage) Mkdir(path string, opt *lib.MkdirOption) error {
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
				if exists, _ := m.Exists(util.NormalizeDir(p)); !exists {
					return m.Put(util.NormalizeDir(p), nil, &store.WriteOptions{IsDir: true})
				}
				return nil
			}
			if err = createParent(filepath.Dir(normalizePath)); err != nil {
				return err
			}
		}
	}
	parentDir := util.NormalizeDir(filepath.Dir(normalizePath))
	if parentDir != "./" {
		if parentExit, _ := m.Exists(parentDir); !parentExit {
			return store.ErrKeyNotFound
		}
	}
	return m.Put(util.NormalizeDir(normalizePath), nil, &store.WriteOptions{
		IsDir: true,
		TTL:   ttl,
	})
}

func (m *EtcdStorage) Rm(path string, recursive bool) error {
	var err error
	if path == "" && recursive {
		kvList, err := m.List(path)
		if err != nil {
			return err
		}
		set := make(map[string]bool)
		for _, p := range kvList {
			pos := strings.IndexRune(p.Key, '/')
			if pos != -1 {
				set[p.Key[:pos]] = true
			} else {
				set[p.Key] = true
			}
		}
		for k, _ := range set {
			if err = m.DeleteTree(k); err != nil {
				return err
			}
		}
	} else {
		err = m.GeneralStorage.Rm(path, recursive)
	}
	return err
}
