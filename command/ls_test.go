package command

import (
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/docker/libkv/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type LsSuite struct {
	suite.Suite
	kv store.Store
}

func (m *LsSuite) SetupSuite() {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	m.kv, err = storage.New(url)
	assert.NoError(m.T(), err)
}

func (m *LsSuite) SetupTest() {
	var err error
	err = m.kv.Put("testroot/", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
	for _, v := range []string{"a", "b"} {
		m.kv.Put("testroot/"+v, nil, &store.WriteOptions{IsDir: false})
	}
	m.kv.Put("testroot/c/", nil, &store.WriteOptions{IsDir: true})
}

func (m *LsSuite) TearDownTest() {
	var err error
	err = m.kv.DeleteTree("testroot/")
	assert.NoError(m.T(), err)
}

func (m *LsSuite) TestListDirectory() {
	var err error
	var paths []string
	ls := LsCommand{}
	pairs, err := ls.ls(m.kv, "testroot")
	for _, v := range pairs {
		paths = append(paths, string(v.Key))
	}
	assert.NoError(m.T(), err)
	assert.NotNil(m.T(), pairs)
	assert.EqualValues(m.T(), []string{"a", "b", "c/"}, paths)
}

func TestRunLsSuite(t *testing.T) {
	s := new(LsSuite)
	suite.Run(t, s)
}
