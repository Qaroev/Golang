package main

import "fmt"

type Planet struct {
	Size   int
	Radius int
}
type World struct {
	Planet
	Name string
}
type IPlanet interface {
	GetSize() int
	SetSize(size int)
}

func (p *Planet) GetSize() int {
	return p.Size * p.Radius
}

func (p *Planet) SetSize(size int) {
	p.Size = size
}
func main() {
	//var mars = &World{Planet{25, 25}, ""}
	//fmt.Println(mars.Radius)
	DoSomething(&World{Planet{25, 25}, "Salom"})
}

func DoSomething(planet IPlanet) {
	fmt.Println(planet.GetSize())
}
