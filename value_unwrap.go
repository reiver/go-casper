package casper

// Unwrap returns the value inside of Value, unless there is nothing inside.
//
// Example
//
//	var value casper.Value
//	
//	// ...
//	
//	str, loaded := value.Unwrap()
//	if !loaded {
//		//@TODO
//	}
func (receiver Value) Unwrap() (string, bool) {
	if NoValue() == receiver {
		return "", false
	}

	return receiver.datum, true
}
