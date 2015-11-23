package lib

import (
	"errors"
	"fmt"
)

var (
	ErrNotSupport     = errors.New("command not support")
	ErrTarAction      = errors.New("unkonw action")
	ErrInputData      = errors.New("need input data")
	ErrArchiveBackend = errors.New("mismatch backend")
	ErrArchiveVersion = errors.New("archive not compatible")
)

func PrefixError(prefix string, err error) string {
	return fmt.Sprintf("%s: %s", prefix, err.Error())
}
