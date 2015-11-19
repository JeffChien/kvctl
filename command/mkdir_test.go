package command

import (
	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/JeffChien/kvctl/lib/storage/etcd"
	"github.com/docker/libkv/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type MkdirSuite struct {
	suite.Suite
	kv store.Store
}

func (m *MkdirSuite) SetupTest() {
	var err error
	err = m.kv.Put("testroot", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
}

func (m *MkdirSuite) TearDownTest() {
	var err error
	err = m.kv.DeleteTree("testroot")
	assert.NoError(m.T(), err)
}

func (m *MkdirSuite) TestNewDirectory() {
	var err error
	cmd, _ := m.kv.(lib.Command)
	err = cmd.Mkdir("testroot/a", nil)
	assert.NoError(m.T(), err)
	err = cmd.Mkdir("testroot/b/c", &lib.MkdirOption{Parent: true})
	assert.NoError(m.T(), err)
	err = cmd.Mkdir("testroot/d", nil)
	assert.NoError(m.T(), err)
	for _, v := range []string{"testroot", "testroot/a", "testroot/b", "testroot/b/c", "testroot/d"} {
		exists, err := m.kv.Exists(v)
		assert.True(m.T(), exists, v)
		assert.NoError(m.T(), err)
	}
}

func TestRunMkdirSuite(t *testing.T) {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "etcd://127.0.0.1:4001"
	}
	s := new(MkdirSuite)
	s.kv, err = storage.New(url)
	assert.NoError(t, err)
	if _, ok := s.kv.(*etcd.EtcdStorage); !ok {
		t.SkipNow()
	}
	suite.Run(t, s)
}
