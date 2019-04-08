package main

import (
	"fmt"
	"net/http"
)

func main() {
	x, _ := http.NewRequest("GET", "https://google.com", nil)
	fmt.Println(x)
}
