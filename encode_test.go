package casper_test

import (
	"github.com/reiver/go-casper"

	"testing"
)

func TestEncodeNilWriter(t *testing.T) {

	var value map[string]interface{} = map[string]interface{}{
		"apple":"one",
		"banana":"two",
		"cherry":"three",
		"something": map[string]string{
			"hello":"world",
			"what":"that",
		},
	}

	if err := casper.Encode(nil, value);  nil == err {
		t.Errorf("Expected an error, but did not actually get one: %#v", err)
		return
	}
}
