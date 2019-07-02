package casper

import (
	"bytes"
	"encoding"
	"fmt"
	"strings"

	"testing"
)

func TestEncodeValueError(t *testing.T) {

	tests := []struct{
		Source interface{}
	}{
		{
			Source: uint64(42),
		},



		{
			Source: float64(123.456789),
		},



		{
			Source: ErrValuer{},
		},



		{
			Source: ErrTextMarshaler{},
		},



		{
			Source: ErrStringCaster{},
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		if err := encodeValue(&buffer, test.Source); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: %#v", testNumber, err)
			t.Errorf("\tSOURCE: (%T) %#v", test.Source, test.Source)
			t.Errorf("\tACTUAL: %q", buffer.String())
			continue
		}
	}
}

func TestEncodeValueNilBytes(t *testing.T) {

	var source []byte = []byte(nil)

	var buffer bytes.Buffer

	if err := encodeValue(&buffer, source); nil != err {
		expected, actual := "", buffer.String()

		t.Errorf("Expect an error, but actually got one: (%T) %q", err, err)
		t.Errorf("\tEXPECTED: %q", expected)
		t.Errorf("\tACTUAL:   %q", actual)
		return
	}

	if expected, actual := "", buffer.String(); expected != actual {
		t.Errorf("\tEXPECTED: %q", expected)
		t.Errorf("\tACTUAL:   %q", actual)
		return
	}
}


func TestEncodeValue(t *testing.T) {

	tests := []struct{
		String   string
		Expected string
	}{
		{
			String:   "",
			Expected: "",
		},



		{
			String:   "Apple",
			Expected: "Apple",
		},
		{
			String:   "Apple banana",
			Expected: "Apple banana",
		},
		{
			String:   "Apple banana CHERRY",
			Expected: "Apple banana CHERRY",
		},



		{
			String:   "nothing()",
			Expected: "nothing()",
		},



		{
			String:   "something(3)",
			Expected: "something(3)",
		},



		{
			String:   `something("x = 1 ; print(x)")`,
			Expected: `something("x = 1 \; print(x)")`,
		},



		{
			String:
`x := 5,
y := 7,
z := z + y,

printlnf("z = %d", z)

for i, e := Elements() {
	printlnf("[%d] ⇒ %v", i, e)
}
`,
			Expected:
`x := 5,
y := 7,
z := z + y,

printlnf("z = %d", z)

for i, e := Elements() {
	printlnf("[%d] ⇒ %v", i, e)
}
`,
		},



		{
			String:
`x := 5;
y := 7;
z := z + y;

printlnf("z = %d", z);

for i, e := Elements() {
	printlnf("[%d] ⇒ %v", i, e);
}
`,
			Expected:
`x := 5\;
y := 7\;
z := z + y\;

printlnf("z = %d", z)\;

for i, e := Elements() {
	printlnf("[%d] ⇒ %v", i, e)\;
}
`,
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		if err := encodeValue(&buffer, test.String); nil != err {
			expected, actual := test.Expected, buffer.String()

			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		var p []byte = []byte(test.String)

		if err := encodeValue(&buffer, p); nil != err {
			expected, actual := test.Expected, buffer.String()

			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		var textMarshaler encoding.TextMarshaler = ToUpper{test.String}

		if err := encodeValue(&buffer, textMarshaler); nil != err {
			expected, actual := test.Expected, buffer.String()

			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}

		if expected, actual := strings.ToUpper(test.Expected), buffer.String(); expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		var stringer fmt.Stringer = ToLower{test.String}

		if err := encodeValue(&buffer, stringer); nil != err {
			expected, actual := test.Expected, buffer.String()

			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}

		if expected, actual := strings.ToLower(test.Expected), buffer.String(); expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		const prefix string = "=====])>"
		var stringcaster interface{ String()(string, error) } = Prefix{Prefix:prefix, Value:test.String}

		if err := encodeValue(&buffer, stringcaster); nil != err {
			expected, actual := test.Expected, buffer.String()

			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}

		if expected, actual := prefix+test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tSTRING:   %q", test.String)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}
}

