package entity

import (
	"errors"
)

var (
	ErrValidation = errors.New("validation")
	ErrNotFound   = errors.New("not found")
	ErrInternal   = errors.New("internal")
	ErrTimeout    = errors.New("timeout")
)
