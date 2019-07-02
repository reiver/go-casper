package casper

import (
	"github.com/reiver/go-whitespace"

	"bytes"
	"io"
	"unicode/utf8"
)

func encodeKey(writer io.Writer, key ...string) error {
	if nil == writer {
		return errNilWriter
	}

	var buffer bytes.Buffer

	for i, token := range key {
		if 0 < i {
			buffer.WriteRune('/')
		}

		var s string = token
		for {
			r, size := utf8.DecodeRuneInString(s)
			if r == utf8.RuneError && 0 == size {
				break
			}
			if r == utf8.RuneError {
				return errRuneError
			}
			if 1 > size {
				return errInternalError
			}

			switch {
			case '/' == r:
				buffer.WriteString(`\/`)
			case ':' == r:
				buffer.WriteString(`\:`)
			case '\\' == r:
				buffer.WriteString(`\\`)
			case '{' == r:
				buffer.WriteString(`\{`)
			case whitespace.IsWhitespace(r):
				buffer.WriteRune('\\')
				buffer.WriteRune(r)
			default:
				buffer.WriteRune(r)
			}

			s = s[size:]
			if 1 > len(s) {
				break
			}
		}
	}

	if _, err := buffer.WriteTo(writer); nil != err {
		return err
	}

	return nil
}
