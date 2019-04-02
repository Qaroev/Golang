package main

import "strconv"

type overLoad struct {
	Name  string
	Age   int
	Wives []string
}

func NewOverLoad(arg interface{}) interface{} {
	o := &overLoad{}
	if val, ok := arg.(int); ok {
		o.Age = val
	} else if val, ok := arg.(string); ok {
		o.Name = val
	} else if val, ok := arg.([]string); ok {
		o.Wives = val
	}
	return o
}

const defaultLen = 10

type nodecha struct {
	Name string
}
type nodes []*nodecha

type num int

func (n num) ToString() string {
	return strconv.FormatInt(int64(n), 10)
}

func (ns nodes) add(n *nodecha) {
	ns = append(ns, n)
}
func (ns nodes) get(i int) *nodecha {
	return ns[i]
}
func NewNodes(val interface{}) nodes {
	if v, ok := val.(int); ok {
		return make(nodes, v)
	}
	if v, ok := val.(nodes); ok {
		n := make(nodes, len(v))
		copy(n, v)
		return n
	} else {
		return make(nodes, defaultLen)
	}
}
