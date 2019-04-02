package main

import "fmt"

type Number struct {
	name map[int]interface{}
}

func (n Number) Summa(property int) bool {
	if _, ok := n.name[property]; ok {
		return true
	} else {
		return false
	}
}

func Delete() {

}

type String struct {
	name map[int]interface{}
}

func (s String) Summa(property int) bool {
	if _, ok := s.name[property]; ok {
		return true
	} else {
		return false
	}
}

type INumber interface {
	Summa(property int) bool
}

func main() {
	n := &Number{}
	n.name = make(map[int]interface{})
	n.name[0] = "Ali"
	n.name[1] = "Vali"
	res := n.Summa(1)
	fmt.Println(n)
	delete(n.name, 0)
	fmt.Println(res)
	fmt.Println(n)

	a := []int{1}
	fmt.Println(a)
	fmt.Println("Len", len(a))
	result := remove(a, len(a)-1)
	fmt.Println(result)
	// [Hello1 Hello2 Hello3]
	a = append(a[:len(a)-1], a[:]...)
	fmt.Println(a)
	fmt.Println(a[:len(a)-1])
	fmt.Println(a[0 : len(a)-1])
	overRide(1, 2, 3, 4, 6)
	b := []float64{1, 2}
	fmt.Println(arrayM(b))
	n1 := &Number{}
	n1.name = make(map[int]interface{})
	n1.name[0] = "Ali"
	arr := []*Number{n, n, n}
	IndexOf(arr, n)

}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func overRide(tab ...int) {
	var ab []int
	for _, k := range tab {
		ab = append(ab, k)
	}
	fmt.Println(ab)
}

func arrayM(column interface{}) bool {
	if _, ok := column.([]float64); ok {
		return true
	} else {
		return false
	}
}

func IndexOf(arr []*Number, str *Number) int {
	//for _, k := range arr {
	//	if k == str {
	//		return true
	//	} else {
	//		return false
	//	}
	//}
	//return false
	for i := 0; i < len(arr); i++ {
		if str == arr[i] {
             return i
		}
	}
	return -1
}
