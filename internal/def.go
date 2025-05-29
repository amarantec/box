package internal

import "errors"

const (
	ZERO    = 0
	EMPTY   = ""
	ENVFILE = ".env"
)

var (
	ErrBookNotFound = errors.New("Book not found")
)
