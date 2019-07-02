package casper

func (receiver Value) Else(datum string) Value {
	if NoValue() != receiver {
		return receiver
	}

	return SomeValue(datum)
}
