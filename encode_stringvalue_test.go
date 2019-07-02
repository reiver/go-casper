package casper

import (
	"bytes"

	"testing"
)

func TestEncodeStringValue(t *testing.T) {

	tests := []struct{
		String  string
		Expected string
	}{
		{
			String:   "",
			Expected: "",
		},



		{
			String:   "apple",
			Expected: "apple",
		},
		{
			String:   "apple banana",
			Expected: "apple banana",
		},
		{
			String:   "apple banana cherry",
			Expected: "apple banana cherry",
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

for i, e := elements() {
	printlnf("[%d] ⇒ %v", i, e)
}
`,
			Expected:
`x := 5,
y := 7,
z := z + y,

printlnf("z = %d", z)

for i, e := elements() {
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

for i, e := elements() {
	printlnf("[%d] ⇒ %v", i, e);
}
`,
			Expected:
`x := 5\;
y := 7\;
z := z + y\;

printlnf("z = %d", z)\;

for i, e := elements() {
	printlnf("[%d] ⇒ %v", i, e)\;
}
`,
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		if err := encodeStringValue(&buffer, test.String); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}
}
