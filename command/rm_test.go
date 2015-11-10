package command

import (
	"github.com/docker/libkv/store"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type RmSuite struct {
	suite.Suite
	kv store.Store
}

func (m *RmSuite) SetupSuite() {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	m.kv, err = storage.NewBackend(url)
	assert.NoError(m.T(), err)
}

func (m *RmSuite) SetupTest() {
	var err error
	m.kv.Put("testroot/", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
	err = m.kv.Put("testroot/validKey", []byte("hello world"), nil)
	assert.NoError(m.T(), err)
	m.kv.Put("testroot/adir/", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
}

func (m *RmSuite) TearDownTest() {
	var err error
	err = m.kv.DeleteTree("testroot/")
	assert.NoError(m.T(), err)
}

func (m *RmSuite) TestRmKey() {
	var err error
	rm := RmCommand{}
	err = rm.rm(m.kv, "testroot/validKey", false)
	assert.NoError(m.T(), err)
}

func (m *RmSuite) TestRmDirWithoutRecursive() {
	var err error
	rm := RmCommand{}
	err = rm.rm(m.kv, "testroot/adir/", false)
	assert.Error(m.T(), err)
}

func (m *RmSuite) TestRmDirWithRecursive() {
	var err error
	rm := RmCommand{}
	err = rm.rm(m.kv, "testroot/adir/", true)
	assert.NoError(m.T(), err)
}

func TestRunRmSuite(t *testing.T) {
	s := new(RmSuite)
	suite.Run(t, s)
}
