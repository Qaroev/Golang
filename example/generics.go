package main

import "errors"

func Add(a, b interface{}) interface{} {
	if ai, ok := a.(int); ok {
		if bi, ok := b.(int); ok {
			return ai + bi
		}else {
			return errors.New("types mismatch")
		}
	}
	if ai, ok := a.(string); ok {
		if bi, ok := b.(string); ok {
			return ai + bi
		}else {
			return errors.New("types mismatch")
		}
	}
	return nil
}
