package casper_test

import (
	"github.com/reiver/go-casper"

	"testing"
)

func TestValueElseUnwrap(t *testing.T) {

	tests := []struct{
		Value    casper.Value
		Else     string
		Expected string
	}{
		{
			Value:    casper.NoValue(),
			Else:                      "defaulted",
			Expected:                  "defaulted",
		},
		{
			Value:    casper.SomeValue("Hi!"),
			Else:                      "defaulted",
			Expected:                  "Hi!",
		},



		{
			Value:    casper.NoValue(),
			Else:                      "Hello world!",
			Expected:                  "Hello world!",
		},
		{
			Value:    casper.SomeValue("Hi!"),
			Else:                      "Hello world!",
			Expected:                  "Hi!",
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.ElseUnwrap(test.Else)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tVALUE: %#v", test.Value)
			t.Errorf("\tELSE:                   %q", test.Else)
			t.Errorf("\tEXPECTED:               %q", expected)
			t.Errorf("\tACTUAL:                 %q", actual)
			continue
		}
	}
}
