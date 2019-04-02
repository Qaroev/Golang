package main

import (
	"github.com/cheekybits/genny/generic"
	"fmt"
)

// Item the type of the Set
type Item generic.Type

// ItemSet the set of Items
type ItemSet struct {
	items map[Item]bool
}

// Add adds a new element to the Set. Returns the Set.
func (s *ItemSet) Add(t Item) ItemSet {
	if s.items == nil {
		s.items = make(map[Item]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return *s
}

// Clear removes all elements from the Set
func (s *ItemSet) Clear() {
	(*s).items = make(map[Item]bool)
}

func main() {
	it := Item("Hello world")
	i := ItemSet{}
	i.items = map[Item]bool{}
	res := i.Add(it)
	fmt.Println(res)
}

