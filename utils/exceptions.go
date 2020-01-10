package utils

import (
	"errors"
)

func Throw(msg string) error {
	return errors.New(msg)
}
