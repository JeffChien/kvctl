package command

import (
	"os"
	"testing"

	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/JeffChien/kvctl/lib/storage/etcd"
	"github.com/docker/libkv/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CatSuite struct {
	suite.Suite
	kv store.Store
}

func (m *CatSuite) SetupTest() {
	var err error
	err = m.kv.Put("testroot", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
	err = m.kv.Put("testroot/validKey", []byte("hello world"), nil)
	assert.NoError(m.T(), err)
}

func (m *CatSuite) TearDownTest() {
	var err error
	err = m.kv.DeleteTree("testroot")
	assert.NoError(m.T(), err)
}

func (m *CatSuite) TestCatValidKey() {
	cmd, _ := m.kv.(lib.Command)
	pair, err := cmd.Cat("testroot/validKey")
	assert.NoError(m.T(), err)
	assert.NotNil(m.T(), pair)
	assert.Equal(m.T(), string(pair.Value), "hello world")
}

func (m *CatSuite) TestCatInvalidKey() {
	cmd, _ := m.kv.(lib.Command)
	pair, err := cmd.Cat("testroot/inValidKey")
	assert.Error(m.T(), err)
	assert.Nil(m.T(), pair)
}

func (m *CatSuite) TestCatDirectory() {
	if _, ok := m.kv.(*etcd.EtcdStorage); !ok {
		m.T().SkipNow()
	}
	cmd, _ := m.kv.(lib.Command)
	pair, err := cmd.Cat("testroot/invalidDir")
	assert.Error(m.T(), err)
	assert.Nil(m.T(), pair)
}

func TestRunCatSuite(t *testing.T) {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	s := new(CatSuite)
	s.kv, err = storage.New(url)
	assert.NoError(t, err)
	suite.Run(t, s)
}
