package order

import (
	"errors"
)

var (
	ErrInvalidValue       = errors.New("invalid value")
	ErrOrdersNotAvailable = errors.New("orders not available")
)
