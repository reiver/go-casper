package casper_test

import (
	"github.com/reiver/go-casper"

	"bytes"
	"fmt"
	"io"
)

func ExampleEncode_mapStringBytes() {

	var data map[string][]byte = map[string][]byte{
		"apple":  []byte("ONE"),
		"banana": []byte("TWO"),
		"cherry": []byte("THREE"),
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
	// apple:	ONE;
	// banana:	TWO;
	// cherry:	THREE;
}
