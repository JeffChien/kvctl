package command

import (
	"fmt"
	"github.com/docker/libkv/store"
)

type ErrorKeyNotFound struct {
	key string
}

func (m *ErrorKeyNotFound) Error() string {
	return fmt.Sprintf("%s: %s", m.key, store.ErrKeyNotFound)
}
