package main

import (
	"fmt"
	"io"
	"os"
)

type phoneReader string

func (p phoneReader) Read(bs []byte) (int, error) {
	count := 0
	for i := 0; i < len(p); i++ {
		if (p[i] >= '0' && p[i] <= '9') {
			bs[count] = p[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	phone1 := phoneReader("+1(234)567 90-10")
	io.Copy(os.Stdout, phone1)
	fmt.Println()
}
