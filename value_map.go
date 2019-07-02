package casper

func (receiver Value) Map(fn func(string)string) Value {
	if NoValue() == receiver {
		return receiver
	}

	return Value {
		loaded: true,
		datum:  fn(receiver.datum),
	}
}
