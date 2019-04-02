package main

import (
	"fmt"
	"errors"
)

type Slice []interface{}

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s *Slice) Add(elem interface{}) error {
	for _, v := range *s {
		if v == elem {
			return errors.New("")
		}
	}
	*s = append(*s, elem)
	return nil
}

func (s *Slice) Remove(elem interface{}) error {
	found := false
	for i, v := range *s {
		if v == elem {
			if i == len(*s)-1 {
				*s = (*s)[:i]
			} else {
				*s = append((*s)[:i], (*s)[i+1:]...)
			}
			found = true
			break
		}
	}
	if !found {
		return errors.New("")
	}
	return nil
}

type student struct {
	id   string
	name string
}

func intSliceExec() {
	slice := NewSlice()
	fmt.Println("current int slice:", slice)
}

func stringSliceExec() {
	slice := NewSlice()
	slice.Add("hello")
}

func structSliceExec() {
	xiaoMing := student{"1001", "xiao ming"}
	slice := NewSlice()
	slice.Add(xiaoMing)
}

func main() {
	intSliceExec()
	stringSliceExec()
	structSliceExec()
}