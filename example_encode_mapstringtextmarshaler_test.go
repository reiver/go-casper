package casper_test

import (
	"github.com/reiver/go-casper"

	"bytes"
	"encoding"
	"fmt"
	"io"
	"strings"
)

type Upper struct {
	value string
}

func (receiver Upper) MarshalText() ([]byte, error) {
	var s string = strings.ToUpper(receiver.value)

	return []byte(s), nil
}

func ExampleEncode_mapStringTextMarshaler() {

	var data map[string]encoding.TextMarshaler = map[string]encoding.TextMarshaler{
		"apple":  Upper{"one"},
		"banana": Upper{"two"},
		"cherry": Upper{"three"},
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
