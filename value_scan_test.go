package casper_test

import (
	"github.com/reiver/go-casper"

	"testing"
)

func TestValueScan(t *testing.T) {

	tests := []struct{
		Datum     interface{}
		Expected  casper.Value
	}{
		{
			Datum:    casper.NoValue(),
			Expected: casper.NoValue(),
		},



		{
			Datum:    casper.SomeValue("Hi!"),
			Expected: casper.SomeValue("Hi!"),
		},
		{
			Datum:    casper.SomeValue("Hello world!"),
			Expected: casper.SomeValue("Hello world!"),
		},
		{
			Datum:    casper.SomeValue("Apple banana CHERRY 🙂 “👾”."),
			Expected: casper.SomeValue("Apple banana CHERRY 🙂 “👾”."),
		},



		{
			Datum:                     "Hi!",
			Expected: casper.SomeValue("Hi!"),
		},
		{
			Datum:                     "Hello world!",
			Expected: casper.SomeValue("Hello world!"),
		},
		{
			Datum:                     "Apple banana CHERRY 🙂 “👾”.",
			Expected: casper.SomeValue("Apple banana CHERRY 🙂 “👾”."),
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

		err := actual.Scan(test.Datum)

		if expected := test.Expected; nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tDatum: (%T) %#v", test.Datum, test.Datum)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tDatum: (%T) %#v", test.Datum, test.Datum)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
