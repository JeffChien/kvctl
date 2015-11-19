package command

import (
	"os"
	"testing"

	"github.com/JeffChien/kvctl/lib"
	"github.com/JeffChien/kvctl/lib/storage"
	"github.com/docker/libkv/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LsSuite struct {
	suite.Suite
	kv store.Store
}

func (m *LsSuite) SetupTest() {
	var err error
	err = m.kv.Put("testroot/", nil, &store.WriteOptions{IsDir: true})
	assert.NoError(m.T(), err)
	for _, v := range []string{"a", "b"} {
		m.kv.Put("testroot/"+v, nil, &store.WriteOptions{IsDir: false})
	}
}

func (m *LsSuite) TearDownTest() {
	var err error
	err = m.kv.DeleteTree("testroot/")
	assert.NoError(m.T(), err)
}

func (m *LsSuite) TestListDirectory() {
	var err error
	var paths []string
	cmd, _ := m.kv.(lib.Command)
	pairs, err := cmd.Ls("testroot/")
	for _, v := range pairs {
		paths = append(paths, string(v.Key))
	}
	assert.NoError(m.T(), err)
	assert.NotNil(m.T(), pairs)
	assert.EqualValues(m.T(), []string{"testroot/a", "testroot/b"}, paths)
}

func TestRunLsSuite(t *testing.T) {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	s := new(LsSuite)
	s.kv, err = storage.New(url)
	assert.NoError(t, err)
	suite.Run(t, s)
}
