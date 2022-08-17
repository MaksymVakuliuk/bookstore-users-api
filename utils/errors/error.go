package errors

import (
	"errors"
)

func NewError(msg string) error {
	return errors.New(msg)
}
