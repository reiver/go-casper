package casper_test

import (
	"github.com/reiver/go-casper"

	"testing"
)

func TestValueUnwrap(t *testing.T) {

	tests := []struct{
		Value          casper.Value
		ExpectedLoaded bool
		Expected       string
	}{
		{
			Value: casper.NoValue(),
			ExpectedLoaded: false,
			Expected:               "",
		},



		{
			Value: casper.SomeValue("Hi!"),
			ExpectedLoaded: true,
			Expected:               "Hi!",
		},



		{
			Value: casper.SomeValue("Hello world!"),
			ExpectedLoaded: true,
			Expected:               "Hello world!",
		},



		{
			Value: casper.SomeValue("Apple banana CHERRY 🙂 “👾”."),
			ExpectedLoaded: true,
			Expected:               "Apple banana CHERRY 🙂 “👾”.",
		},
	}

	for testNumber, test := range tests {

		actual, actualLoaded := test.Value.Unwrap()

		if expected, expectedLoaded := test.Expected, test.ExpectedLoaded; expectedLoaded != actualLoaded {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED loaded: %t", expectedLoaded)
			t.Errorf("\tACTUAL   loaded: %t", actualLoaded)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if expected, expectedLoaded := test.Expected, test.ExpectedLoaded; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED loaded: %t", expectedLoaded)
			t.Errorf("\tACTUAL   loaded: %t", actualLoaded)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
