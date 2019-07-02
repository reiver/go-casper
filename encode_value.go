package casper

import (
	"encoding"
	"fmt"
	"io"
	"reflect"
	"unsafe"
)

func encodeValue(writer io.Writer, value interface{}) error {
	if nil == writer {
		return errNilWriter
	}

	if nil == value {
		return nil
	}

	var stringValue string

	switch casted := value.(type) {
	case Valuer:
		s, err := casted.CASPERValue()
		if nil != err {
			return err
		}

		stringValue = s
	case encoding.TextMarshaler:
		p, err := casted.MarshalText()
		if nil != err {
			return err
		}

		bh := (*reflect.SliceHeader)(unsafe.Pointer(&p))

		sh := reflect.StringHeader{
			Data: bh.Data,
			Len: bh.Len,
		}

		stringValue = *(*string)(unsafe.Pointer(&sh))
	case StringCaster:
		s, err := casted.String()
		if nil != err {
			return err
		}

		stringValue = s
	case fmt.Stringer:
		stringValue = casted.String()
	case []byte:
		bh := (*reflect.SliceHeader)(unsafe.Pointer(&casted))

		sh := reflect.StringHeader{
			Data: bh.Data,
			Len: bh.Len,
		}

		stringValue = *(*string)(unsafe.Pointer(&sh))
	case string:
		stringValue = casted
	default:
		return unsupportedSource(value)
	}

	return encodeStringValue(writer, stringValue)
}
