package casper

import (
	"database/sql/driver"
)

// Value makes casper.Value fit the database/sql/driver.Valuer interface.
func (receiver Value) Value() (driver.Value, error) {
	if NoValue() == receiver {
		return receiver, errNotLoaded
	}

	return receiver.datum, nil
}
