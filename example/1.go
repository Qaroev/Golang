package main

import (
	"fmt"
	"strings"
)

func main() {

	//s := make([]int, 0)
	//s = append(s, 1, 2, 3)
	//for len(s) > 0 {
	//	s = append(s[:0], s[1:]...)
	//	i := len(s) - 1
	//	ss := s[i]
	//	fmt.Println(s)
	//	fmt.Println(ss)
	//
	//}

	names := strings.Fields("Hello world")
	fmt.Println(names)
}
