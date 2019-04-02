package main


var IString string
type Number struct {
	name string
}

func NewNumber(name string) Number {
	n := Number{}
	n.name = name
	return n
}
