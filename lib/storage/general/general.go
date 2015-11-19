package general

import (
	"github.com/JeffChien/kvctl/lib"
	"github.com/docker/libkv/store"
)

type GeneralStorage struct {
	store.Store
}

func New(s store.Store) *GeneralStorage {
	return &GeneralStorage{Store: s}
}

func (m *GeneralStorage) Cat(path string) (*store.KVPair, error) {
	return m.Get(path)
}

func (m *GeneralStorage) Ls(path string) ([]*store.KVPair, error) {
	return m.List(path)
}

func (m *GeneralStorage) Mkdir(path string, opt *lib.MkdirOption) error {
	return lib.ErrNotSupport
}

func (m *GeneralStorage) Rm(path string, recursive bool) error {
	var err error
	if recursive {
		err = m.DeleteTree(path)
	} else {
		err = m.Delete(path)
	}
	return err
}

func (m *GeneralStorage) Touch(path string, data []byte, opts *store.WriteOptions) error {
	return m.Put(path, data, opts)
}
