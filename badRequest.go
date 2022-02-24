package errors

import (
	"errors"
	"fmt"
)

var (
	errBadRequest = errors.New("Bad Request")
)

func NewBadRequest(msg string) error {
	return fmt.Errorf("%s: %w", msg, errBadRequest)
}

func IsBadRequest(err error) bool {
	return errors.Is(err, errBadRequest)
}
