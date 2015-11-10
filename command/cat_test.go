package command

import (
	"github.com/docker/libkv/store"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type CatSuite struct {
	suite.Suite
	kv store.Store
}

func (m *CatSuite) SetupSuite() {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	m.kv, err = storage.NewBackend(url)
	assert.NoError(m.T(), err)
}

func (m *CatSuite) SetupTest() {
	var err error
	err = m.kv.Put("testroot/", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
	err = m.kv.Put("testroot/validKey", []byte("hello world"), nil)
	assert.NoError(m.T(), err)
}

func (m *CatSuite) TearDownTest() {
	var err error
	err = m.kv.DeleteTree("testroot/")
	assert.NoError(m.T(), err)
}

func (m *CatSuite) TestCatValidKey() {
	cat := CatCommand{}
	pair, err := cat.cat(m.kv, "testroot/validKey")
	assert.NoError(m.T(), err)
	assert.NotNil(m.T(), pair)
	assert.Equal(m.T(), string(pair.Value), "hello world")
}

func (m *CatSuite) TestCatInvalidKey() {
	cat := CatCommand{}
	pair, err := cat.cat(m.kv, "testroot/inValidKey")
	assert.Error(m.T(), err)
	assert.Nil(m.T(), pair)
}

func (m *CatSuite) TestCatDirectory() {
	cat := CatCommand{}
	pair, err := cat.cat(m.kv, "testroot/invalidDir/")
	assert.Error(m.T(), err)
	assert.Nil(m.T(), pair)
}

func TestRunCatSuite(t *testing.T) {
	s := new(CatSuite)
	suite.Run(t, s)
}
