package main

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
)

func main() {
	buf := [3]byte{}
	size := 0
	var err error

	size, err = push(buf[:], size, 'a')
	fmt.Printf("buf=%v max=%d size=%d err=[%v]\n", buf, len(buf), size, err)
	size, err = push(buf[:], size, 'b')
	fmt.Printf("buf=%v max=%d size=%d err=[%v]\n", buf, len(buf), size, err)
	size, err = push(buf[:], size, 'c')
	fmt.Printf("buf=%v max=%d size=%d err=[%v]\n", buf, len(buf), size, err)
	size, err = push(buf[:], size, 'd')
	fmt.Printf("buf=%v max=%d size=%d err=[%v]\n", buf, len(buf), size, err)
}

func push(buf []byte, size int, b byte) (int, error) {
	max := len(buf)

	if max < 1 {
		return size, fmt.Errorf("buffer underflow: max=%d char=%d", max, b)
	}

	if size >= max {
		return size, fmt.Errorf("buffer overflow: size=%d max=%d char=%d", size, max, b)
	}

	buf[size] = b

	return size + 1, nil
}

var typse generic.Type



