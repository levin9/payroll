package blls

import (
	"errors"
)

func throwException(msg string) error {
	return errors.New(msg)
}
