package errors

import (
	"errors"
	"fmt"
)

var (
	errNotFound = errors.New("Not Found")
)

func NewNotFound(msg string) error {
	return fmt.Errorf("%s: %w", msg, errNotFound)
}

func IsNotFound(err error) bool {
	return errors.Is(err, errNotFound)
}
