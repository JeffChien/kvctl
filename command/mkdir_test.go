package command

import (
	"github.com/docker/libkv/store"
	"github.com/jeffchien/kvctl/lib/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type MkdirSuite struct {
	suite.Suite
	kv store.Store
}

func (m *MkdirSuite) SetupSuite() {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	m.kv, err = storage.NewBackend(url)
	assert.NoError(m.T(), err)
}

func (m *MkdirSuite) SetupTest() {
	var err error
	err = m.kv.Put("testroot/", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
}

func (m *MkdirSuite) TearDownTest() {
	var err error
	err = m.kv.DeleteTree("testroot/")
	assert.NoError(m.T(), err)
}

func (m *MkdirSuite) TestNewDirectory() {
	var err error
	mkdir := MkdirCommand{}
	err = mkdir.mkdir(m.kv, "testroot/a", nil)
	assert.NoError(m.T(), err)
	err = mkdir.mkdir(m.kv, "testroot/b/c", &mkdirOption{Parent: true})
	assert.NoError(m.T(), err)
	err = mkdir.mkdir(m.kv, "testroot/d/", nil)
	assert.NoError(m.T(), err)
	for _, v := range []string{"testroot/", "testroot/a/", "testroot/b/", "testroot/b/c/", "testroot/d/"} {
		exists, err := m.kv.Exists(v)
		assert.True(m.T(), exists, v)
		assert.NoError(m.T(), err)
	}
}

func TestRunMkdirSuite(t *testing.T) {
	s := new(MkdirSuite)
	suite.Run(t, s)
}
