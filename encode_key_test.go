package casper

import (
	"bytes"

	"testing"
)

func TestEncodeKey(t *testing.T) {

	tests := []struct{
		Key      []string
		Expected string
	}{
		{
			Key: []string{""},
			Expected:     "",
		},



		{
			Key: []string{`\`},
			Expected:     `\\`,
		},



		{
			Key: []string{"A"},
			Expected:     "A",
		},
		{
			Key: []string{"B"},
			Expected:     "B",
		},
		{
			Key: []string{"C"},
			Expected:     "C",
		},



		{
			Key: []string{"a"},
			Expected:     "a",
		},
		{
			Key: []string{"b"},
			Expected:     "b",
		},
		{
			Key: []string{"c"},
			Expected:     "c",
		},



		{
			Key: []string{"\u0009"}, // horizontal tab
			Expected:   "\\\u0009",
		},
		{
			Key: []string{"\u000A"}, // line feed
			Expected:   "\\\u000A",
		},
		{
			Key: []string{"\u000B"}, // vertical tab
			Expected:   "\\\u000B",
		},
		{
			Key: []string{"\u000C"}, // form feed
			Expected:   "\\\u000C",
		},
		{
			Key: []string{"\u000D"}, // carriage return
			Expected:   "\\\u000D",
		},
		{
			Key: []string{"\u0020"}, // space
			Expected:   "\\\u0020",
		},
		{
			Key: []string{"\u0085"}, // next line
			Expected:   "\\\u0085",
		},
		{
			Key: []string{"\u00A0"}, // no-break space
			Expected:   "\\\u00A0",
		},
		{
			Key: []string{"\u1680"}, // ogham space mark
			Expected:   "\\\u1680",
		},
		{
			Key: []string{"\u180E"}, // mongolian vowel separator
			Expected:   "\\\u180E",
		},
		{
			Key: []string{"\u2000"}, // en quad
			Expected:   "\\\u2000",
		},
		{
			Key: []string{"\u2001"}, // em quad
			Expected:   "\\\u2001",
		},
		{
			Key: []string{"\u2002"}, // en space
			Expected:   "\\\u2002",
		},
		{
			Key: []string{"\u2003"}, // em space
			Expected:   "\\\u2003",
		},
		{
			Key: []string{"\u2004"}, // three-per-em space
			Expected:   "\\\u2004",
		},
		{
			Key: []string{"\u2005"}, // four-per-em space
			Expected:   "\\\u2005",
		},
		{
			Key: []string{"\u2006"}, // six-per-em space
			Expected:   "\\\u2006",
		},
		{
			Key: []string{"\u2007"}, // figure space
			Expected:   "\\\u2007",
		},
		{
			Key: []string{"\u2008"}, // punctuation space
			Expected:   "\\\u2008",
		},
		{
			Key: []string{"\u2009"}, // thin space
			Expected:   "\\\u2009",
		},
		{
			Key: []string{"\u200A"}, // hair space
			Expected:   "\\\u200A",
		},
		{
			Key: []string{"\u2028"}, // line separator
			Expected:   "\\\u2028",
		},
		{
			Key: []string{"\u2029"}, // paragraph separator
			Expected:   "\\\u2029",
		},
		{
			Key: []string{"\u202F"}, // narrow no-break space
			Expected:   "\\\u202F",
		},
		{
			Key: []string{"\u205F"}, // medium mathematical space
			Expected:   "\\\u205F",
		},
		{
			Key: []string{"\u3000"}, // ideographic space
			Expected:   "\\\u3000",
		},



		{
			Key: []string{"Hello\u0009world!"}, // horizontal tab
			Expected:   "Hello\\\u0009world!",
		},
		{
			Key: []string{"Hello\u000Aworld!"}, // line feed
			Expected:   "Hello\\\u000Aworld!",
		},
		{
			Key: []string{"Hello\u000Bworld!"}, // vertical tab
			Expected:   "Hello\\\u000Bworld!",
		},
		{
			Key: []string{"Hello\u000Cworld!"}, // form feed
			Expected:   "Hello\\\u000Cworld!",
		},
		{
			Key: []string{"Hello\u000Dworld!"}, // carriage return
			Expected:   "Hello\\\u000Dworld!",
		},
		{
			Key: []string{"Hello\u0020world!"}, // space
			Expected:   "Hello\\\u0020world!",
		},
		{
			Key: []string{"Hello\u0085world!"}, // next line
			Expected:   "Hello\\\u0085world!",
		},
		{
			Key: []string{"Hello\u00A0world!"}, // no-break space
			Expected:   "Hello\\\u00A0world!",
		},
		{
			Key: []string{"Hello\u1680world!"}, // ogham space mark
			Expected:   "Hello\\\u1680world!",
		},
		{
			Key: []string{"Hello\u180Eworld!"}, // mongolian vowel separator
			Expected:   "Hello\\\u180Eworld!",
		},
		{
			Key: []string{"Hello\u2000world!"}, // en quad
			Expected:   "Hello\\\u2000world!",
		},
		{
			Key: []string{"Hello\u2001world!"}, // em quad
			Expected:   "Hello\\\u2001world!",
		},
		{
			Key: []string{"Hello\u2002world!"}, // en space
			Expected:   "Hello\\\u2002world!",
		},
		{
			Key: []string{"Hello\u2003world!"}, // em space
			Expected:   "Hello\\\u2003world!",
		},
		{
			Key: []string{"Hello\u2004world!"}, // three-per-em space
			Expected:   "Hello\\\u2004world!",
		},
		{
			Key: []string{"Hello\u2005world!"}, // four-per-em space
			Expected:   "Hello\\\u2005world!",
		},
		{
			Key: []string{"Hello\u2006world!"}, // six-per-em space
			Expected:   "Hello\\\u2006world!",
		},
		{
			Key: []string{"Hello\u2007world!"}, // figure space
			Expected:   "Hello\\\u2007world!",
		},
		{
			Key: []string{"Hello\u2008world!"}, // punctuation space
			Expected:   "Hello\\\u2008world!",
		},
		{
			Key: []string{"Hello\u2009world!"}, // thin space
			Expected:   "Hello\\\u2009world!",
		},
		{
			Key: []string{"Hello\u200Aworld!"}, // hair space
			Expected:   "Hello\\\u200Aworld!",
		},
		{
			Key: []string{"Hello\u2028world!"}, // line separator
			Expected:   "Hello\\\u2028world!",
		},
		{
			Key: []string{"Hello\u2029world!"}, // paragraph separator
			Expected:   "Hello\\\u2029world!",
		},
		{
			Key: []string{"Hello\u202Fworld!"}, // narrow no-break space
			Expected:   "Hello\\\u202Fworld!",
		},
		{
			Key: []string{"Hello\u205Fworld!"}, // medium mathematical space
			Expected:   "Hello\\\u205Fworld!",
		},
		{
			Key: []string{"Hello\u3000world!"}, // ideographic space
			Expected:   "Hello\\\u3000world!",
		},



		{
			Key: []string{"apple"},
			Expected:     "apple",
		},



		{
			Key: []string{"apple","banana"},
			Expected:     "apple/banana",
		},
		{
			Key: []string{"apple/banana"},
			Expected:     `apple\/banana`,
		},



		{
			Key: []string{"apple","banana","cherry"},
			Expected:     "apple/banana/cherry",
		},
		{
			Key: []string{"apple/banana","cherry"},
			Expected:     `apple\/banana/cherry`,
		},
		{
			Key: []string{"apple/banana/cherry"},
			Expected:     `apple\/banana\/cherry`,
		},





		{
			Key: []string{"ONE", "two", "Three"},
			Expected:     `ONE/two/Three`,
		},



		{
			Key: []string{"🙂", "😍😁", "👽👾👻"},
			Expected:     `🙂/😍😁/👽👾👻`,
		},



		{
			Key: []string{"apple", "banana", "\u0009", "cherry", "Hello\u3000world!", "C:\\Documents\\Vacation", "\\👾"},
			Expected:     "apple/banana/\\\u0009/cherry/Hello\\\u3000world!/C\\:\\\\Documents\\\\Vacation/\\\\👾",
		},
	}

	for testNumber, test := range tests {
		var buffer bytes.Buffer

		if err := encodeKey(&buffer, test.Key...); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d, the encoded key is not what was expected.", testNumber)
			t.Logf("EXPECTED: ...\n%s", expected)
			t.Logf("ACTUAL: ...\n%s", actual)
			continue
		}
	}
}
