package casper

import (
	"fmt"
)

// GoString makes casper.Value fit the fmt.GoStringer interface.
//
// It gets used with the %#v verb with the printing family of functions
// in the Go built-in "fmt" package.
//
// I.e., it gets used with: fmt.Fprint(), fmt.Fprintf(), fmt.Fprintln(),
// fmt.Print(), fmt.Printf(), fmt.Println(), fmt.Sprint(), fmt.Sprintf(),
// fmt.Sprintln().
//
// Example
//
//	var value casper.Value
//	
//	// ...
//	
//	fmt.Printf("value = %#v\n", value) // <---- value.GoString() is called by fmt.Printf()
func (receiver Value) GoString() string {
	if NoValue() == receiver {
		return "casper.NoValue()"
	}

	return fmt.Sprintf("casper.SomeValue(%q)", receiver.datum)
}
