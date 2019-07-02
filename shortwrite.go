package casper

import (
	"fmt"
)

// ShortWrite is the error returned from casper.Encode() when, for whatever reason, not everything
// was able to be written to an io.Writer.
//
// For example:
//
//	err := casper.Encode(writer, source)
//	
//	if nil != err {
//		switch err.(type) {
//		case casper.ShortWrite:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type ShortWrite interface {
	error

	ShortWrite()

	// Returns what trying to be written.
	Source() string

	// Returns the number of bytes that were expected to be written.
	Expected() int64

	// Returns the number of bytes that were actually written.
	Actual() int64
}

type internalShortWrite struct {
	source   string
	expected int64
	actual   int64
}

func (receiver internalShortWrite) Error() string {
	return fmt.Sprintf("casper: ushort write: expected to write %d bytes, but actually wrote %d bytes.", receiver.expected, receiver.actual)
}

func (receiver internalShortWrite) Source() string {
	return receiver.source
}

func (receiver internalShortWrite) Expected() int64 {
	return receiver.expected
}

func (receiver internalShortWrite) Actual() int64 {
	return receiver.actual
}

func (internalShortWrite) ShortWrite() {
	// Nothing here.
}
