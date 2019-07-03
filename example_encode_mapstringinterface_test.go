package casper_test

import (
	"github.com/reiver/go-casper"

	"bytes"
	"fmt"
	"io"
)

func ExampleEncode_mapStringInterface() {

	var data map[string]interface{} = map[string]interface{}{
		"apple":  "one",
		"banana": "two",
		"cherry": "three",

		"http": map[string]interface{}{
			"port": "8080",
			"log": map[string]interface{}{
				"dest":"stdout",
			},
		},

		"x-notice": Measurement{3, "cm"},
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
	// http {
	// 	log {
	// 		dest:	stdout;
	// 	}
	// 	port:	8080;
	// }
	// x-notice:	3cm;
}
