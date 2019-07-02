package casper

type StringCaster interface{
	String()(string, error)
}
