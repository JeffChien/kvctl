package lib

import (
	"errors"
	"fmt"
)

var (
	ErrNotSupport = errors.New("command not support")
)

func PrefixError(prefix string, err error) string {
	return fmt.Sprintf("%s: %s", prefix, err.Error())
}
