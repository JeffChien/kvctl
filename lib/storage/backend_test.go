package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidUrl(t *testing.T) {
	assert := assert.New(t)
	_, err := NewBackend("I'm not a valid url")
	assert.NotNil(err)
}
