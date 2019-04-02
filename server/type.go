package main

import "fmt"

type mile int

func main() {

	var distance mile = 5
	fmt.Println(distance)
	distance += 5
	fmt.Println(distance)
	fn()
}

type library []string

func printBooks(lib library) {

	for _, value := range lib {

		fmt.Println(value)
	}
}
func fn() {
	var myLibrary library = library{"Book1", "Book2", "Book3"}
	printBooks(myLibrary)
}
