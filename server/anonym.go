package main

import "fmt"

func selectFn(n int) (func(int, int) int) {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}

type array struct {
	adadi  int
	adadi1 int
}

type arrrr []interface{}

func main() {
	res := arrrr{array{1, 3}, array{1, 3}}
	ress := Contains(res, 2)
	fmt.Println(ress)
	f := selectFn(1)
	fmt.Println(f(2, 3))
	fmt.Println(f(4, 5))
	fmt.Println(f(7, 6))
}
func Contains(arr []interface{}, str interface{}) bool {
	for _, ar := range arr {
		if ar == str {
			return true
		} else {
			return false
		}
	}
	return false
}
