package casper

import (
	"fmt"
	"runtime"
	"strings"
)

// UnsupportedSource is the error returned from casper.Encode(), and casper.Value.Scan()
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

	Debug() string
	Source() interface{}
}

type internalUnsupportedSource struct {
	source interface{}
	file   string
	line   int
}

func unsupportedSource(source interface{}) UnsupportedSource {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		line = -1
	}

	return internalUnsupportedSource{
		source: source,
		file:   file,
		line:   line,
	}
}

func (receiver internalUnsupportedSource) Debug() (s string) {
	defer func() {
		if r := recover(); nil != r {
			s = receiver.Error()
		}
	}()

	var file string = receiver.file
	{
		var index int = strings.LastIndex(file, "/")
		if 0 < index {
			file = file[1+index:]
		}
	}

	return fmt.Sprintf("casper: %s:%d: unsupported source: %T", file, receiver.line, receiver.source)
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
