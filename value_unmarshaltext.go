package casper

// Scan makes casper.Value fit the encoding.TextUnmarshaler interface.
func (receiver *Value) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.Scan(text)
}
