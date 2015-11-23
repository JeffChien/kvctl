package lib

import (
	"github.com/docker/libkv/store"
	"time"
)

type MkdirOption struct {
	Parent bool
	TTL    time.Duration
}

type Command interface {
	Cat(path string) (*store.KVPair, error)
	Ls(path string) ([]*store.KVPair, error)
	Mkdir(path string, opt *MkdirOption) error
	Rm(path string, recursive bool) error
	Touch(path string, data []byte, opts *store.WriteOptions) error
	Dump(path string) ([]byte, error)
	Restore(archive []byte) error
}
