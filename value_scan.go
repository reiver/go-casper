package casper

// Scan makes casper.Value fit the database/sql.Scanner interface.
func (receiver *Value) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch casted := src.(type) {
	case string:
		*receiver = SomeValue(casted)
		return nil
	case []byte:
		switch casted {
		case nil:
			*receiver = NoValue()
		default:
			*receiver = SomeValue(string(casted))
		}
		return nil
	case Value:
		*receiver = casted
		return nil
	default:
		return internalUnsupportedSource{src}
	}
}
