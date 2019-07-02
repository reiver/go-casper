package casper

func (receiver Value) CASPERValue() (string, error) {
	return receiver.String()
}
