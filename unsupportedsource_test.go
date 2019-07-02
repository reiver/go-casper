package casper

import (
	"testing"
)

func TestInternalUnsupportedSourceAsError(t *testing.T) {
	var err error = internalUnsupportedSource{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == err {
		t.Errorf("This should never happen.")
		return
	}
}
