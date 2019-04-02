package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Simple shell")
	//fmt.Println("--------------------------")
	//
	//for {
	//	fmt.Print("-->")
	//	text, _ := reader.ReadString('\n')
	//	text = strings.Replace(text, "\n", "", -1)
	//	if strings.Compare("hi", text) == 0 {
	//		fmt.Println("Hello YouSelf")
	//		break
	//	}
	//}
	fmt.Println(User{}.Name)
	Scanner()
}

func Scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("--------------------------")
	for scanner.Scan() {
		fmt.Print("-->")
		fmt.Println(scanner.Text())
	}
}

type User struct {
	Name bool "true"
}
