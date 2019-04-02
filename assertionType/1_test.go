package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestNewAccount(t *testing.T) {
	var s = types("hello world")
	var s1 = types(25)
	res := NewAccount(s, 12)
	res1 := NewAccount(s1, 12)
	stringType := reflect.TypeOf(res).Name()
	numType := reflect.TypeOf(res1).Name()
	if "string" != stringType {
		fmt.Println("Error this is not type string")
	}
	if "number" != numType {
		fmt.Println("Error this is not type number")
	}
}
