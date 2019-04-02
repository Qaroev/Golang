package main

type person struct {
	name     string
	username string
	age      int
}

func (r *person) isFirst() {
	r.username = "Qaroev"
	r.name = "Gulboy"
	r.age = 22
}


func main()  {
	ma1()
}

func (r *person) ma1() {
	r.isFirst()
}