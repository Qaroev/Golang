package main

import "fmt"

func main() {
	var x = 4
	var p *int
	p = &x
	fmt.Println(p)
	fnNew()
}
func fnNew() {
	p := new(int)
	fmt.Println("Value:", *p) // Value: 0 - значение по умолчанию
	*p = 8                    // изменяем значение
	fmt.Println("Value:", *p) // Value: 8
	newChange()
}
func changeValue(x *int) {
	*x = (*x) * (*x)
}
func newChange() {

	d := 5
	fmt.Println("d before:", d) // 5
	changeValue(&d)             // изменяем значение
	fmt.Println("d after:", d)  // 25 - значение изменилось!
}
