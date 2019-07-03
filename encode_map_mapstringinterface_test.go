package casper

import (
	"bytes"

	"testing"
)

func TestEncodeMapStringInterface(t *testing.T) {

	tests := []struct{
		Source   map[string]interface{}
		Expected string
	}{
		{
			Source: map[string]interface{}{},
			Expected: "",
		},



		{
			Source: map[string]interface{}{
				"apple":"one",
			},
			Expected: "apple:\tone;\n",
		},
		{
			Source: map[string]interface{}{
				"apple":"one",
				"banana":"two",
			},
			Expected: "apple:\tone;\nbanana:\ttwo;\n",
		},
		{
			Source: map[string]interface{}{
				"apple":"one",
				"banana":"two",
				"cherry":"three",
			},
			Expected:
`apple:	one;
banana:	two;
cherry:	three;
`,
		},



		{
			Source: map[string]interface{}{
				"apple":"one",
				"banana":"two",
				"cherry":"three",

				"melons": map[string]string{
					"ananas":"four",
					"cantaloupe":"five",
					"honeydew":"six",
					"watermelon":"seven",
				},
			},
			Expected:
`apple:	one;
banana:	two;
cherry:	three;
melons	{
	ananas:	four;
	cantaloupe:	five;
	honeydew:	six;
	watermelon:	seven;
}
`,
		},



		{
			Source: map[string]interface{}{
				"apple":"one",
				"banana":"two",
				"cherry":"three",

				"melons": map[string]interface{}{
					"ananas":"four",
					"cantaloupe":"five",
					"honeydew":"six",
					"watermelon":"seven",

					"notes": map[string]string{
						"ONE":"Hello",
						"TWO":"world",
					},
				},
			},
			Expected:
`apple:	one;
banana:	two;
cherry:	three;
melons	{
	ananas:	four;
	cantaloupe:	five;
	honeydew:	six;
	notes	{
		ONE:	Hello;
		TWO:	world;
	}
	watermelon:	seven;
}
`,
		},



		{
			Source: map[string]interface{}{
				"👾":"one",
				"two":"😈😠",
				"😱😨😮":"three",
			},
			Expected:
				"two:\t😈😠;\n"+
				"👾:\tone;\n"+
				"😱😨😮:\tthree;\n",
		},
		{
			Source: map[string]interface{}{
				"👾":"one",
				"two":"😈😠",
				"😱😨😮":"three",

				"🎃🎃🎃":map[string]interface{}{
					"01":"💀",
					"02":"👻👻",
					"03":"🦇🦇🦇",
					"04":map[string]string{
						"a":"A",
						"b":"B",
						"c":"C",
					},
				},
			},
			Expected:
				"two:\t😈😠;\n"+
				"🎃🎃🎃\t{\n"+
					"\t01:\t💀;\n"+
					"\t02:\t👻👻;\n"+
					"\t03:\t🦇🦇🦇;\n"+
					"\t04\t{\n"+
						"\t\ta:\tA;\n"+
						"\t\tb:\tB;\n"+
						"\t\tc:\tC;\n"+
					"\t}\n"+
				"}\n"+
				"👾:\tone;\n"+
				"😱😨😮:\tthree;\n",
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
			t.Errorf("For test #%d, the encoded value is not what was expected.", testNumber)
			t.Logf("\tSOURCE: %#v", test.Source)
			t.Logf("\tEXPECTED: ...\n%s", expected)
			t.Logf("\tACTUAL: ...\n%s", actual)
			continue
		}
	}
}
