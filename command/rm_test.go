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

type RmSuite struct {
	suite.Suite
	kv store.Store
}

func (m *RmSuite) SetupTest() {
	var err error
	m.kv.Put("testroot", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
	err = m.kv.Put("testroot/validKey", []byte("hello world"), nil)
	assert.NoError(m.T(), err)
	m.kv.Put("testroot/adir", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
}

func (m *RmSuite) TearDownTest() {
	var err error
	err = m.kv.DeleteTree("testroot")
	assert.NoError(m.T(), err)
}

func (m *RmSuite) TestRmKey() {
	var err error
	cmd, _ := m.kv.(lib.Command)
	err = cmd.Rm("testroot/validKey", false)
	assert.NoError(m.T(), err)
}

func (m *RmSuite) TestRmDirWithoutRecursive() {
	var err error
	cmd, _ := m.kv.(lib.Command)
	err = cmd.Rm("testroot/adir", false)
	switch m.kv.(type) {
	case *etcd.EtcdStorage:
		assert.Error(m.T(), err)
	default:
		assert.NoError(m.T(), err)
	}
}

func (m *RmSuite) TestRmDirWithRecursive() {
	var err error
	cmd, _ := m.kv.(lib.Command)
	err = cmd.Rm("testroot/adir", true)
	assert.NoError(m.T(), err)
}

func (m *RmSuite) TestRmRootRecursive() {
	var err error
	cmd, _ := m.kv.(lib.Command)
	err = cmd.Rm("", true)
	assert.NoError(m.T(), err)
	err = m.kv.Put("testroot", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
}

func TestRunRmSuite(t *testing.T) {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	s := new(RmSuite)
	s.kv, err = storage.New(url)
	assert.NoError(t, err)
	suite.Run(t, s)
}
