package casper

// Value represents a CAPSER value, in the key-value sense of the word “value”.
//
// It is NOT required the this casper.Value type be used to use the tools in this package.
//
// This package works fine without this.
//
// This is provided as a convenience for those looking for a string ‘option type’, that also
// works seemlessly with the built-in "database/sql", "encoding", "fmt", and "encoding/json" packges.
type Value struct {
	datum  string
	loaded bool
}

func NoValue() Value {
	return Value{}
}

func SomeValue(value string) Value {
	return Value{
		loaded: true,
		datum:  value,
	}
}
