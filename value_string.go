package casper

func (receiver Value) String() (string, error) {
	if NoValue() == receiver {
		return "", errNotLoaded
	}

	return receiver.datum, nil
}
