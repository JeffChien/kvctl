package command

import (
	"github.com/docker/libkv/store"
	"github.com/jeffchien/kvctl/lib/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type TouchSuite struct {
	suite.Suite
	kv store.Store
}

func (m *TouchSuite) SetupSuite() {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	m.kv, err = storage.NewBackend(url)
	assert.NoError(m.T(), err)
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
	touch := TouchCommand{}
	err := touch.touch(m.kv, "testroot/validKey", []byte("hello"), nil)
	assert.NoError(m.T(), err)
	pair, err := m.kv.Get("testroot/validKey")
	assert.NoError(m.T(), err)
	assert.NotNil(m.T(), pair)
	assert.Equal(m.T(), string(pair.Value), "hello")
}

func (m *TouchSuite) TestTouchDirectoryIsNotAllowed() {
	touch := TouchCommand{}
	err := touch.touch(m.kv, "testroot/invalidDir/", []byte("hello"), nil)
	assert.Error(m.T(), err)
}

func (m *TouchSuite) TestTouchKeyWhileParentNotExist() {
	touch := TouchCommand{}
	err := touch.touch(m.kv, "testroot/invalidDir/validKey", []byte("hello"), nil)
	assert.Error(m.T(), err)
}

func TestRunTouchSuite(t *testing.T) {
	s := new(TouchSuite)
	suite.Run(t, s)
}
