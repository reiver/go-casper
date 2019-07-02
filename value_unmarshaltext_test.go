package casper_test

import (
	"github.com/reiver/go-casper"

	"testing"
)

func TestValueTextUnmarshaler(t *testing.T) {

	tests := []struct{
		Datum   []byte
		Expected  casper.Value
	}{
		{
			Datum:              []byte(nil),
			Expected: casper.NoValue(),
		},



		{
			Datum:              []byte("Hi!"),
			Expected: casper.SomeValue("Hi!"),
		},
		{
			Datum:              []byte("Hello world!"),
			Expected: casper.SomeValue("Hello world!"),
		},
		{
			Datum:              []byte("Apple banana CHERRY 🙂 “👾”."),
			Expected: casper.SomeValue("Apple banana CHERRY 🙂 “👾”."),
		},
	}

	for testNumber, test := range tests {

		var actual casper.Value

		err := actual.UnmarshalText(test.Datum)

		if expected := test.Expected; nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tDATUM: (%T) %#v", test.Datum, test.Datum)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tDATUM: (%T) %#v", test.Datum, test.Datum)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
