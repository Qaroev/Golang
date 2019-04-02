package main

import "testing"

//func TestNewOverLoad(t *testing.T) {
//	o := NewOverLoad("Bobojon")
//	if o.Name != "Bobojon" {
//		t.Fail()
//	}
//	o = NewOverLoad(20)
//	if o.Age != 20 {
//		t.Fail()
//	}
//	o = NewOverLoad([]string{"Sitora", "safdasf", "asfsf"})
//	if o.Wives[0] != "Sitora" {
//		t.Fail()
//	}
//}

func TestNum_ToString(t *testing.T) {
	n := num(1)
	val := n.ToString()
	if val != "1" {
		t.Fail()
	}

}
func TestNum_ToString2(t *testing.T) {
	o := NewNodes(1)
	if o != nil {
		t.Fail()
	}
	n := nodes{}
	o1 := NewNodes(n)
	if o1 != nil {
		t.Fail()
	}
	n1 := nodes{}
	n2 := &nodecha{"Bobojon"}
	n1.add(n2)
	if n2.Name != "Bobojon" {
		t.Fail()
	}
	n6 := nodes{}
	n5 := &nodecha{"Bobojon"}
	n11 := &nodecha{"Bobojon"}
	n6.add(n11)
	n6.add(n5)
	re := n6.get(0)
	if re.Name != "Bobojon" {
		t.Fail()
	}

}
