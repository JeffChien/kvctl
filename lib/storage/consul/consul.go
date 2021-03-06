package consul

import (
	"github.com/JeffChien/kvctl/lib/storage/general"
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"
)

type ConsulStorage struct {
	*general.GeneralStorage
}

func New(hosts []string, config *store.Config) (store.Store, error) {
	consul.Register()

	s, err := libkv.NewStore(
		store.CONSUL,
		hosts,
		config,
	)
	if err != nil {
		return nil, err
	}
	return &ConsulStorage{
		GeneralStorage: general.New(s, string(store.CONSUL)),
	}, nil
}
