package casper_test

import (
	"github.com/reiver/go-casper"

	"bytes"
	"fmt"
	"io"
	"strings"
)

type Title struct {
	value string
}

func (receiver Title) String() (string, error) {
	var s string = strings.Title(receiver.value)

	return s, nil
}

func ExampleEncode_mapStringStringCaster() {

	var data map[string]casper.StringCaster = map[string]casper.StringCaster{
		"apple":  Title{"one"},
		"banana": Title{"two"},
		"cherry": Title{"three"},
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
	// apple:	One;
	// banana:	Two;
	// cherry:	Three;
}
