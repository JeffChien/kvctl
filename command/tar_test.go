package command

import (
	"os"
	"testing"

	//"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	//"github.com/JeffChien/kvctl/lib/storage/etcd"
	"github.com/docker/libkv/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TarSuite struct {
	suite.Suite
	kv store.Store
}

//TODO: add test
func (m *TarSuite) SetupTest() {
}

func (m *TarSuite) TearDownTest() {
}

func TestRunTarSuite(t *testing.T) {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	s := new(TarSuite)
	s.kv, err = storage.New(url)
	assert.NoError(t, err)
	suite.Run(t, s)
}
