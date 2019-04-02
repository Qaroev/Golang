package main

import (
	"strings"
	"math/rand"
	"fmt"
)

const (
	bufferSize     = 0x20000
	unicodeBom     = 0xFEFF
	boundaryLength = 32
)

var mimeBoundaryChars = []byte("-_1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	res := mimeBoundary()
	fmt.Println(res)
}

func mimeBoundary() string {
	mime := strings.Builder{}
	for i := 0; i < boundaryLength; i++ {
		s := rand.Intn(int(len(mimeBoundaryChars)))
		mime.WriteRune(rune(mimeBoundaryChars[s]))
	}
	return mime.String()
}
