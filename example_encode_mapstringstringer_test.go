package casper_test

import (
	"github.com/reiver/go-casper"

	"bytes"
	"fmt"
	"io"
)

type Optional struct {
	loaded bool
	value string
}

func (receiver Optional) String() string {
	if !receiver.loaded {
		return "Nothing()"
	}

	return fmt.Sprintf("Something(%q)", receiver.value)
}

func Nothing() Optional {
	return Optional{}
}

func Something(v string) Optional {
	return Optional{
		loaded: true,
		value: v,
	}
}

func ExampleEncode_mapStringStringer() {

	var data map[string]fmt.Stringer = map[string]fmt.Stringer{
		"apple":  Something("red"),
		"banana": Something("yellow"),
		"cherry": Nothing(),
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
	// apple:	Something("red");
	// banana:	Something("yellow");
	// cherry:	Nothing();
}
