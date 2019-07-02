package casper

func (receiver Value) Then(fn func(string)Value) Value {
	if NoValue() == receiver {
		return receiver
	}

	return fn(receiver.datum)
}
