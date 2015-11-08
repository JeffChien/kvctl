package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type BackendSuite struct {
	suite.Suite
}

func (m *BackendSuite) TestInitialize() {
	var err error
	var url string
	if url = os.Getenv("KVCTL_BACKEND"); url == "" {
		url = "consul://127.0.0.1:8500"
	}
	kv, err := NewBackend(url)
	assert.NotNil(m.T(), kv)
	assert.NoError(m.T(), err)
}

func TestInvalidUrl(t *testing.T) {
	s := new(BackendSuite)
	suite.Run(t, s)
}
