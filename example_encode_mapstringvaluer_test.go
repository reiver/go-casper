package casper_test

import (
	"github.com/reiver/go-casper"

	"bytes"
	"fmt"
	"io"
)

type Measurement struct {
	value int
	unit  string
}

func (receiver Measurement) CASPERValue() (string, error) {
	var buffer bytes.Buffer

	_, err := fmt.Fprintf(&buffer, "%d%s", receiver.value, receiver.unit)
	if nil != err {
		return "", err
	}

	return buffer.String(), nil
}

func ExampleEncode_mapStringValuer() {

	var data map[string]casper.Valuer = map[string]casper.Valuer{
		"apple":  Measurement{5, "cm"},
		"banana": Measurement{7, "in"},
		"cherry": Measurement{12, "lbs"},
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
	// apple:	5cm;
	// banana:	7in;
	// cherry:	12lbs;
}
