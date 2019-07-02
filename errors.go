package casper

import (
	"errors"
)

var (
	errInternalError = errors.New("casper: Internal Error")
	errNilWriter     = errors.New("casper: Nil Writer")
	errRuneError     = errors.New("casper: Rune Error")
)
