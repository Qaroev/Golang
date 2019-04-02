package _package

import "fmt"

func isTrue() {
	n := 0
	if true {
		n = 1
		n++
	}
	fmt.Println(n) // 2
}
