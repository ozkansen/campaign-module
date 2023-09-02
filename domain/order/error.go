package order

import (
	"errors"
)

var (
	ErrInvalidValue       = errors.New("invalid value")
	ErrOrdersNotAvailable = errors.New("err orders not available")
)
