package casper

import (
	"bytes"

	"testing"
)

func TestEncodeMapMapStringString(t *testing.T) {

	tests := []struct{
		Source   map[string]string
		Expected string
	}{
		{
			Source: map[string]string{},
			Expected: "",
		},



		{
			Source: map[string]string{
				"apple":"one",
			},
			Expected: "apple:\tone;\n",
		},
		{
			Source: map[string]string{
				"apple":"one",
				"banana":"two",
			},
			Expected: "apple:\tone;\nbanana:\ttwo;\n",
		},
		{
			Source: map[string]string{
				"apple":"one",
				"banana":"two",
				"cherry":"three",
			},
			Expected: "apple:\tone;\nbanana:\ttwo;\ncherry:\tthree;\n",
		},



		{
			Source: map[string]string{
				"👾":"one",
				"two":"😈😠",
				"😱😨😮":"three",
			},
			Expected: "two:\t😈😠;\n👾:\tone;\n😱😨😮:\tthree;\n",
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		if err := encodeMap(&buffer, test.Source); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Logf("\tSOURCE: %#v", test.Source)
			t.Logf("\tEXPECTED: ...\n%s", test.Expected)
			t.Logf("\tACTUAL BUFFER (so far): ...\n%s", buffer.String())
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Logf("\tSOURCE: %#v", test.Source)
			t.Logf("\tEXPECTED: ...\n%s", expected)
			t.Logf("\tACTUAL: ...\n%s", actual)
			continue
		}
	}
}

