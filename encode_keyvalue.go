package casper

import (
	"io"
)

func encodeKeyValue(writer io.Writer, key string, value interface{}) error {
	if nil == writer {
		return errNilWriter
	}

	if err := encodeKey(writer, key); nil != err {
		return err
	}

	if err := encodeDelimiter(writer); nil != err {
		return err
	}

	if err := encodeTab(writer); nil != err {
		return err
	}

	if err := encodeValue(writer, value); nil != err {
		return err
	}

	if err := encodeSemicolon(writer); nil != err {
		return err
	}

	if err := encodeNewLine(writer); nil != err {
		return err
	}

	return nil
}
