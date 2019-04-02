package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
	fn()
}
func fn() {
	s := make([]int, 50)
	fmt.Println(s)
	c := s[:2]
	fmt.Println(c)
	fn1()
}
func fn1() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	apend()
}
func apend() {
	var s []int

	s = append(s, 0)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	rang()
}
func rang() {
	var pow = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range pow {
		fmt.Printf("2*%d = %d\n", i, v)
	}
}
func ran()  {
	//pow:make()
}
