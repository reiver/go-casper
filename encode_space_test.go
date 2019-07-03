package casper

import (
	"bytes"

	"testing"
)

func TestEncodeSpace(t *testing.T) {

	var buffer bytes.Buffer

	if err := encodeSpace(&buffer); nil != err {
		t.Errorf("Did not expect an error, but actually got one.")
		t.Logf("ERROR type: %T", err)
		t.Logf("ERRORe: %q", err)
		return
	}

	if expected, actual := " ", buffer.String(); expected != actual {
		t.Errorf("Expected a single space to be written, but that is not what actually got.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
