package main

import (
	"fmt"
	"reflect"
)

type IType interface{}

type Accountable interface {
	GetId() IType
	GetSum() int
	SetSum(sum int)
}

type Accounts struct {
	Accountable
	id  IType
	sum int
}

func NewAccounts(id IType, sum int) Accounts {
	a := Accounts{}
	a.id = id
	a.sum = sum
	return a
}

func (a Accounts) GetId() IType {
	return a.id
}

func (a Accounts) GetSum() int {
	return a.sum
}

func (a Accounts) SetSum(sum int) {
	a.sum = sum
}

func main() {
	var ac = NewAccounts(Accounts{nil, "vghkjfhhj", 3}, 5000)
	r := ac.GetId()
	fmt.Println(reflect.TypeOf(r.id))
	fmt.Println(ac.GetSum())
	var name = 4555
	ac.SetSum(name)
	fmt.Println(ac.sum)
}
