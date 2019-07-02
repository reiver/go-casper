package casper

import (
	"errors"
	"fmt"
	"strings"
)


type ToUpper struct {
	Value string
}

func (receiver ToUpper) MarshalText() ([]byte, error) {
	var s string = strings.ToUpper(receiver.Value)

	var p []byte = []byte(s)

	return p, nil
}


type ToLower struct {
	Value string
}

func (receiver ToLower) String() string {
	return strings.ToLower(receiver.Value)
}


type Prefix struct {
	Prefix string
	Value  string
}

func (receiver Prefix) String() (string, error) {
	s := fmt.Sprintf("%s%s", receiver.Prefix, receiver.Value)

	return s, nil
}


type ErrValuer struct{}

func (receiver ErrValuer) CASPERValue() (string, error) {
		return "", errors.New("casper: ErrValuer return an error from .CASPERValue()")
}


type ErrTextMarshaler struct {}

func (receiver ErrTextMarshaler) MarshalText() ([]byte, error) {
		return []byte(nil), errors.New("casper: ErrTextMarshaler return an error from .MarshalText()")
}


type ErrStringCaster struct {}

func (receiver ErrStringCaster) String() (string, error) {
		return "", errors.New("casper: ErrStringCaster return an error from .String()")
}
