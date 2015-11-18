package storage

import (
	boltdbStorage "github.com/JeffChien/kvctl/lib/storage/boltdb"
	consulStorage "github.com/JeffChien/kvctl/lib/storage/consul"
	etcdStorage "github.com/JeffChien/kvctl/lib/storage/etcd"
	zookeeperStorage "github.com/JeffChien/kvctl/lib/storage/zookeeper"
	"github.com/docker/libkv/store"
	"net/url"
	"time"
)

func New(backend string) (store.Store, error) {
	var storage store.Backend
	var factory func(hosts []string, config *store.Config) (store.Store, error)
	u, err := url.Parse(backend)
	if err != nil {
		return nil, err
	}

	storage = store.Backend(u.Scheme)
	switch storage {
	case store.CONSUL:
		factory = consulStorage.New
	case store.ETCD:
		factory = etcdStorage.New
	case store.ZK:
		factory = zookeeperStorage.New
	case store.BOLTDB:
		factory = boltdbStorage.New
	default:
		storage = "unknow"
	}
	return factory([]string{u.Host}, &store.Config{
		ConnectionTimeout: 10 * time.Second,
	})
}
