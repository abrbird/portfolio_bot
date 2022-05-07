package domain

import (
	"errors"
)

var (
	UnknownError       = errors.New("unknown error")
	ErrorNotFound      = errors.New("not found error")
	ErrorAlreadyExists = errors.New("already exists error")
)
