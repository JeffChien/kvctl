package command

import (
	"fmt"
)

func PrefixError(prefix string, err error) string {
	return fmt.Sprintf("%s: %s", prefix, err.Error())
}
