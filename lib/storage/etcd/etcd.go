package etcd

import (
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/etcd"
)

type EtcdStorage struct {
	store.Store
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
		Store: s,
	}, nil
}
