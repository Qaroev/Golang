package main

import (
	"os"
	"fmt"
)

func main() {
	data, err := os.Open("READ_FILE/user.json")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	fmt.Println(*data)
}
