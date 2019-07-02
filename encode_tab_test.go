package casper

import (
	"bytes"

	"testing"
)

func TestEncodeTab(t *testing.T) {

	var buffer bytes.Buffer

	if err := encodeTab(&buffer); nil != err {
		t.Errorf("Did not expect an error, but actually got one.")
		t.Logf("ERROR type: %T", err)
		t.Logf("ERRORe: %q", err)
		return
	}

	if expected, actual := "\t", buffer.String(); expected != actual {
		t.Errorf("Expected a single tab to be written, but that is not what actually got.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
