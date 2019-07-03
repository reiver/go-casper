package casper_test

import (
	"github.com/reiver/go-casper"

	"bytes"
	"fmt"
	"io"
)

func ExampleEncode_mapStringString() {

	var data map[string]string = map[string]string{
		"apple":  "one",
		"banana": "two",
		"cherry": "three",
	}

	var buffer bytes.Buffer

	var writer io.Writer = &buffer

	err := casper.Encode(writer, data)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Print(buffer.String())

	// Output:
	// apple:	one;
	// banana:	two;
	// cherry:	three;
}
