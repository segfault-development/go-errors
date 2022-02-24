package errs

import (
	"errors"
	"fmt"
)

var (
	errInternal = errors.New("Internal")
)

func NewInternal(msg string) error {
	return fmt.Errorf("%s: %w", msg, errInternal)
}

func IsInternal(err error) bool {
	return errors.Is(err, errInternal)
}
