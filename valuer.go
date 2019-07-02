package casper

// If a type has fits the casper.Valuer interface, then casper.Encode() func will use
// its .CASPERValue() method.
type Valuer interface {
	CASPERValue() (string, error)
}
