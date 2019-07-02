package casper_test

import (
	"github.com/reiver/go-casper"

	"testing"
)

func TestValueElse(t *testing.T) {

	tests := []struct{
		Value    casper.Value
		Else     string
		Expected casper.Value
	}{
		{
			Value:    casper.NoValue(),
			Else:                      "defaulted",
			Expected: casper.SomeValue("defaulted"),
		},
		{
			Value:    casper.SomeValue("Hi!"),
			Else:                      "defaulted",
			Expected: casper.SomeValue("Hi!"),
		},



		{
			Value:    casper.NoValue(),
			Else:                      "Hello world!",
			Expected: casper.SomeValue("Hello world!"),
		},
		{
			Value:    casper.SomeValue("Hi!"),
			Else:                      "Hello world!",
			Expected: casper.SomeValue("Hi!"),
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.Else(test.Else)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tVALUE:    %#v", test.Value)
			t.Errorf("\tELSE:                      %q", test.Else)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
