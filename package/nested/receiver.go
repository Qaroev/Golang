package main

import (
	"math"
	"fmt"
)

type velox struct {
	x, y float64
}

func (v velox) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (v *velox) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := velox{3, 4}
	v.scale(10)
	fmt.Println(v.Abs())
	isTrue()
}
func isTrue() {
	var s float64 = 56
	fmt.Printf("type %T\n", s)
	var d = int(s)
	fmt.Printf("type %T\n", d)
	mainFunc()
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

func ScalFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func mainFunc() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScalFunc(&v, 10)
	p := &Vertex{5, 6}
	p.Scale(5)
	ScalFunc(p, 5)
	fmt.Println(v, p)
}
