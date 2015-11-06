package storage

import (
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
	"github.com/docker/libkv/store/consul"
	"github.com/docker/libkv/store/etcd"
	"github.com/docker/libkv/store/zookeeper"
	"net/url"
	"time"
)

func NewBackend(backend string) (store.Store, error) {
	var storage store.Backend
	u, err := url.Parse(backend)
	if err != nil {
		return nil, err
	}

	storage = store.Backend(u.Scheme)

	switch storage {
	case store.CONSUL:
		consul.Register()
	case store.ETCD:
		etcd.Register()
	case store.ZK:
		zookeeper.Register()
	case store.BOLTDB:
		boltdb.Register()
	default:
		storage = "unknow"
	}

	return libkv.NewStore(
		storage,
		[]string{u.Host},
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
}
