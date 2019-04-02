package main

import (
	"fmt"
	"reflect"
)

/*
class Account<T>{

    private T id;
    private int sum;

    Account(T id, int sum){
        this.id = id;
        this.sum = sum;
    }

    public T getId() { return id; }
    public int getSum() { return sum; }
    public void setSum(int sum) { this.sum = sum; }
}
 */

type types interface{}

type Account struct {
	id  types
	sum int
}

func NewAccount(id types, sum int) Account {
	a := Account{}
	a.id = id
	a.sum = sum
	return a
}

func (a Account) GetId() types {
	return a.id
}

func (a Account) GetSum() types {
	return a.sum
}

func (a Account) SetSum(sum int) {
	a.sum = sum
}

func main() {
	var s = types("hello world")
	var s1 = types(25)
	res := NewAccount(s, 12)
	res1 := NewAccount(s1, 12)
	r := res.GetId()
	r1 := res1.GetId()
	fmt.Println(r)
	fmt.Println(r1)
	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.TypeOf(r1))
}
