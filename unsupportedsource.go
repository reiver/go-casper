package casper

import (
	"fmt"
)

// UnsupportedSource is the error returned from capser.Encode(), and casper.Value.Scan()
// when the type of the source is not supported.
//
// For example:
//
//	var source int64 // <---- Note that the type is int64, which casper.Encode() does not support (by itself). So it will return an error.
//	
//	// ...
//	
//	err := casper.Encode(writer, source)
//	
//	if nil != err {
//		switch err.(type) {
//		case casper.UnsupportedSource:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
//
// Also, for example:
//
//	var source int64 // <---- Note that the type is int64, which casper.Encode() does not support (by itself). So it will return an error.
//	
//	// ...
//	
//	var value casper.Value
//
//	// ...
//
//	err := value.Scan(source)
//	
//	if nil != err {
//		switch err.(type) {
//		case casper.UnsupportedSource:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type UnsupportedSource interface {
	error

	UnsupportedSource()

	Soure() interface{}
}

type internalUnsupportedSource struct {
	source interface{}
}

func (receiver internalUnsupportedSource) Error() string {
	return fmt.Sprintf("casper: unsupported source: %T", receiver.source)
}

func (receiver internalUnsupportedSource) Source() interface{} {
	return receiver.source
}

func (internalUnsupportedSource) UnsupportedSource() {
	// Nothing here.
}
