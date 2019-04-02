package main

import (
	"testing"
	"fmt"
)

func TestNegnot(t *testing.T) {
	var i Deha
	a := []Deha{&Negnot{10}, &Panjakent{100}}
	for _, v := range a {
		i = v
		fmt.Println(i.GetDistance())
	}
}
