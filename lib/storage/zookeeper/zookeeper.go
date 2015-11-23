package zookeeper

import (
	"github.com/JeffChien/kvctl/lib/storage/general"
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/zookeeper"
)

type ZookeeperStorage struct {
	*general.GeneralStorage
}

func New(hosts []string, config *store.Config) (store.Store, error) {
	zookeeper.Register()

	s, err := libkv.NewStore(
		store.ZK,
		hosts,
		config,
	)
	if err != nil {
		return nil, err
	}
	return &ZookeeperStorage{
		GeneralStorage: general.New(s, string(store.ZK)),
	}, nil
}
