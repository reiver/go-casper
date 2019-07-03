package casper

import (
	"testing"
)

func TestInternalShortWriteAsError(t *testing.T) {
	var err error = internalShortWrite{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == err {
		t.Errorf("This should never happen.")
		return
	}
}

func TestInternalShortWriteAsShortWrite(t *testing.T) {
	var err ShortWrite = internalShortWrite{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == err {
		t.Errorf("This should never happen.")
		return
	}
}
