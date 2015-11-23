package boltdb

import (
	"github.com/JeffChien/kvctl/lib/storage/general"
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
)

type BoltdbStorage struct {
	*general.GeneralStorage
}

func New(hosts []string, config *store.Config) (store.Store, error) {
	boltdb.Register()

	s, err := libkv.NewStore(
		store.BOLTDB,
		hosts,
		config,
	)
	if err != nil {
		return nil, err
	}
	return &BoltdbStorage{
		GeneralStorage: general.New(s, string(store.BOLTDB)),
	}, nil
}
