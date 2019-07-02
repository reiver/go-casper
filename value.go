package casper

// Value represents a CAPSER value, in the key-value sense of the word “value”.
//
// It is NOT required for this casper.Value type be used to use the tools in this package.
//
// This package works fine without this casper.Value.
//
// casper.Value is provided as a convenience for those looking for a string ‘option type’, that also
// works seemlessly with the built-in "database/sql", "encoding", "fmt", and "encoding/json" packages.
type Value struct {
	datum  string
	loaded bool
}

// NoValue creates a casper.Value with no value.
//
// You can use this in if-statements, and switch-statements too.
//
// Example
//
// Here is an example of using casper.NoValue() with an assignment:
//
//	var value casper.Value = casper.NoValue()
//
// Example
//
// Here is an example of using casper.NoValue() with an if-statement:
//
//	var value casper.Value
//	
//	// ...
//	
//	if casper.NoValue() == value {
//		//@TODO
//	}
//
// Example
//
// Here is an example of using casper.NoValue() with an switch-statement:
//
//	var value casper.Value
//	
//	// ...
//	
//	switch value {
//	case casper.NoValue():
//		//@TODO
//	default:
//		//@TODO
//	}
func NoValue() Value {
	return Value{}
}

// SomeValue creates a casper.Value with a specific value.
//
// You can use this in if-statements, and switch-statements too.
//
// Example
//
// Here is an example of using casper.SomeValue() with an assignment:
//
//	var value casper.Value = casper.SomeValue("Hello world!")
//
// Example
//
// Here is an example of using casper.SomeValue() with an if-statement:
//
//	var value casper.Value
//	
//	// ...
//	
//	if casper.SomeValue("Hello world!") == value {
//		//@TODO
//	}
//
// Example
//
// Here is an example of using casper.SomeValue() with an switch-statement:
//
//	var value casper.Value
//	
//	// ...
//	
//	switch value {
//	case casper.SomeValue("Hello world!"):
//		//@TODO
//	default:
//		//@TODO
//	}
func SomeValue(value string) Value {
	return Value{
		loaded: true,
		datum:  value,
	}
}
