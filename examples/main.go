package main

import (
	"fmt"
	"reflect"
)

func main() {
	var input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	isTrue(input)
	res := int(5.0) % 2.0
	fmt.Println(reflect.TypeOf(res), res)
	IString = "Bobojon"
	res := NewNumber("Qaroev")
	fmt.Println(res, IString)
}

func isTrue(is [][]int) {
	res := is[1:len(is)]
	fmt.Println(res)
}
