package command

import (
	"os"
	"testing"

	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/JeffChien/kvctl/lib/storage/consul"
	"github.com/JeffChien/kvctl/lib/storage/zookeeper"
	"github.com/docker/libkv/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TouchSuite struct {
	suite.Suite
	kv store.Store
}

func (m *TouchSuite) SetupTest() {
	var err error
	err = m.kv.Put("testroot/", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
}

func (m *TouchSuite) TearDownTest() {
	var err error
	err = m.kv.DeleteTree("testroot/")
	assert.NoError(m.T(), err)
}

func (m *TouchSuite) TestTouchKey() {
	cmd, _ := m.kv.(lib.Command)
	err := cmd.Touch("testroot/validKey", []byte("hello"), nil)
	assert.NoError(m.T(), err)
	pair, err := m.kv.Get("testroot/validKey")
	assert.NoError(m.T(), err)
	assert.NotNil(m.T(), pair)
	assert.Equal(m.T(), string(pair.Value), "hello")
}

func (m *TouchSuite) TestTouchKeyWithSlashIsAllowed() {
	if _, ok := m.kv.(*consul.ConsulStorage); !ok {
		m.T().SkipNow()
	}
	cmd, _ := m.kv.(lib.Command)
	err := cmd.Touch("testroot/invalidDir/", []byte("hello"), nil)
	assert.NoError(m.T(), err)
}

func (m *TouchSuite) TestTouchKeyWhileParentNotExist() {
	cmd, _ := m.kv.(lib.Command)
	err := cmd.Touch("testroot/invalidDir/validKey", []byte("hello"), nil)
	switch m.kv.(type) {
	case *zookeeper.ZookeeperStorage:
		assert.Error(m.T(), err)
	default:
		assert.NoError(m.T(), err)
	}
}

func TestRunTouchSuite(t *testing.T) {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	s := new(TouchSuite)
	s.kv, err = storage.New(url)
	assert.NoError(t, err)
	suite.Run(t, s)
}
