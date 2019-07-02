package casper

import (
	"bytes"
	"io"
	"unicode/utf8"
)

func encodeStringValue(writer io.Writer, value string) error {
	if nil == writer {
		return errNilWriter
	}

	var buffer bytes.Buffer
	{
		var s string = value
		for 0 < len(s) {

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

			switch r {
			case ';':
				buffer.WriteString(`\;`)
			default:
				buffer.WriteRune(r)
			}

			s = s[size:]
		}
	}

	bufferLen    := buffer.Len()
	bufferString := buffer.String()

	n, err := buffer.WriteTo(writer);
	if nil != err {
		return err
	}

	if expected, actual := int64(bufferLen), n; expected != actual {
		return internalShortWrite{
			source: bufferString,
			expected: int64(bufferLen),
			actual:   n,
		}
	}

	return nil
}
