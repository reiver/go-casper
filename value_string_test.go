package casper_test

import (
	"github.com/reiver/go-casper"

	"testing"
)

func TestValueString(t *testing.T) {

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

		actual, err := test.Value.String()

		if !test.ExpectedLoaded && nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: %#v", testNumber, err)
			t.Errorf("\tEXPECTED loaded: %t", test.ExpectedLoaded)
			t.Errorf("\tEXPECTED: %#v", test.Expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if test.ExpectedLoaded && nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tEXPECTED loaded: %t", test.ExpectedLoaded)
			t.Errorf("\tEXPECTED: %#v", test.Expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED loaded: %t", test.ExpectedLoaded)
			t.Errorf("\tACTUAL   err:   (%T) %q", err, err)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
