package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	open()
}

func read() {
	data, err := ioutil.ReadFile("EXAMPLE_WORKS/read.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func write() {
	mydata := []byte("We’ll use Go’s while loop equivalent of a for loop without any parameters to ensure our program continues on forever. In this example every time text is entered and then enter is pressed, we assign text to equal everything up to and including the \n special character.")
	err := ioutil.WriteFile("EXAMPLE_WORKS/read.txt", mydata, 0)

	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile("EXAMPLE_WORKS/read.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func open() {

	data, err := os.OpenFile("EXAMPLE_WORKS/read.txt", os.O_APPEND, 0)

	if err != nil {
		panic(err)
	}

	defer data.Close()

	if _, err := data.WriteString("If we want to do a comparison on the string that has just been entered then we can use the strings.Replace method in order to remove this trailing \n character with nothing and then do the comparison."); err != nil {
		panic(err)
	}
	//
	//f, err := ioutil.ReadFile("EXAMPLE_WORKS/read.txt")
	//if err!=nil{
	//	panic(err)
	//}
	//
	//
	//fmt.Println(f)

}
