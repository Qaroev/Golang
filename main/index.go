package main

import (
	"fmt"
	"unsafe"
)

type Deha interface {
	GetDistance() int
}
type Negnot struct {
	distance int
}

type Panjakent struct {
	distance int
}

func (p *Panjakent) GetDistance() int {
	return p.distance
}
func (p *Negnot) GetDistance() int {
	return p.distance
}

func main() {
	bar := Deha(&Panjakent{100})
	i, ok := bar.(*Panjakent)
	i.distance = 1000
	ng := Negnot{10}
	fmt.Println(i.GetDistance(), ok)
	ng.incnAssertion(ng)
	incnAssertionNoCheck(ng)
}

func (p *Negnot) incnAssertion(any Negnot) {
	res := *(**Panjakent)(unsafe.Pointer(&any))
	if res != nil {
		res.GetDistance()
	}
}

func incnAssertionNoCheck(any Negnot) int {
	res := *(**Panjakent)(unsafe.Pointer(&any))
	return res.GetDistance()
}
