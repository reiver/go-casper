package casper

import (
	"encoding"
	"fmt"
	"io"
)

// Encode encodes ‘src’ as a Casper file data format, and writes it to ‘writer’.
//
// ‘src’ can be either any of the following map types:
//
// • map[string]casper.Valuer
//
// • map[string]encoding.TextMarshaler
//
// • map[string]fmt.StringCaster
//
// • map[string]fmt.Stringer
//
// • map[string][]byte
//
// • map[string]string
//
// • map[string]interface{}
//
// And can also be a struct that the appropriate “casper” struct tag.
func Encode(writer io.Writer, src interface{}) error {
	if nil == writer {
		return errNilWriter
	}

	if nil == src {
		return nil
	}

	switch casted := src.(type) {

	case map[string]Valuer:
		return encodeMap(writer, casted)

	case map[string]encoding.TextMarshaler:
		return encodeMap(writer, casted)

	case map[string]StringCaster:
		return encodeMap(writer, casted)

	case map[string]fmt.Stringer:
		return encodeMap(writer, casted)

	case map[string][]byte:
		return encodeMap(writer, casted)

	case map[string]string:
		return encodeMap(writer, casted)

	case map[string]interface{}:
		return encodeMap(writer, casted)

	default:
//@TODO: Struct, with tags

		return unsupportedSource(src)
	}
}
