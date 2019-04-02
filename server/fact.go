package main

import "fmt"

func factorial(n int) int{

	if n == 0{
		return 1
	}
	return n * factorial(n - 1)
}
func main() {

	fmt.Println(factorial(4))   // 24
	fmt.Println(factorial(5))   // 120
	fmt.Println(factorial(6))   // 720
}