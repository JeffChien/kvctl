package boltdb

import (
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
)

type BoltdbStorage struct {
	store.Store
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
		Store: s,
	}, nil
}
