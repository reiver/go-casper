package casper

import (
	"errors"
)

var (
	errInternalError = errors.New("casper: Internal Error")
	errNilReceiver   = errors.New("casper: Nil Receiver")
	errNilWriter     = errors.New("casper: Nil Writer")
	errNotLoaded     = errors.New("casper: Not Loaded")
	errRuneError     = errors.New("casper: Rune Error")
)
