package main

import "fmt"

func fibbonachi(n uint) uint{

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n - 1) + fibbonachi(n - 2)
}
func main() {

	fmt.Println(fibbonachi(4))
	fmt.Println(fibbonachi(5))
	fmt.Println(fibbonachi(6))
}