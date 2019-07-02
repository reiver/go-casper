package casper_test

import (
	"github.com/reiver/go-casper"

	"testing"
)

func TestValueMap(t *testing.T) {

	tests := []struct{
		Value    casper.Value
		Func     func(string)string
		Expected casper.Value
	}{
		{
			Value: casper.NoValue(),
			Func:  func(string)string{
				return "Hello world!"
			},
			Expected: casper.NoValue(),
		},
		{
			Value: casper.SomeValue("Hi!"),
			Func:  func(string)string{
				return "Hello world!"
			},
			Expected: casper.SomeValue("Hello world!"),
		},



		{
			Value: casper.NoValue(),
			Func:  func(s string)string{
				return "«" + s + "»"
			},
			Expected: casper.NoValue(),
		},
		{
			Value: casper.SomeValue("What?"),
			Func:  func(s string)string{
				return "Joe said, «" + s + "»"
			},
			Expected: casper.SomeValue("Joe said, «What?»"),
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.Map(test.Func)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
