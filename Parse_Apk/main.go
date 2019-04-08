package main

import (
	"fmt"
	"github.com/phinexdaz/ipapk"
)

func main() {
	apk, _ := ipapk.NewAppParser("Parse_Apk/test/36.apk")
	fmt.Println(apk)
}
