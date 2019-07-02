package casper

func (receiver Value) ElseUnwrap(datum string) string {
	if NoValue() == receiver {
		return datum
	}

	return receiver.datum
}
