package product

import (
	"errors"
)

var (
	ErrInvalidValue         = errors.New("invalid value")
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrProductNotFound      = errors.New("product not found")
)
