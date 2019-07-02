package casper

import (
	"io"
)

func encodeTab(writer io.Writer) error {
	if nil == writer {
		return errNilWriter
	}

	var buffer [1]byte = [1]byte{'\t'}

	n, err := writer.Write(buffer[:])
	if nil != err {
		return err
	}

	if expected, actual := len(buffer), n; expected != actual {
		return io.ErrShortWrite
	}

	return nil
}

