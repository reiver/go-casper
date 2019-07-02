package casper_test

import (
	"github.com/reiver/go-casper"

	"testing"
)

func TestValueGoString(t *testing.T) {

	tests := []struct{
		Value    casper.Value
		Expected string
	}{
		{
			Value:     casper.NoValue(),
			Expected: "casper.NoValue()",
		},



		{
			Value:     casper.SomeValue("Hi!"),
			Expected: `casper.SomeValue("Hi!")`,
		},



		{
			Value:     casper.SomeValue("Hello world!"),
			Expected: `casper.SomeValue("Hello world!")`,
		},



		{
			Value:     casper.SomeValue("Apple banana CHERRY 🙂 “👾”."),
			Expected: `casper.SomeValue("Apple banana CHERRY 🙂 “👾”.")`,
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.GoString()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
