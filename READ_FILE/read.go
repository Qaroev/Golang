package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	open()
}

func write() {
	mydata := []byte("All the data I wish to write to a fileOnce we have constructed this byte array, we can then call ioutil.WriteFile() to write this byte array to a file. The WriteFile() method takes in 3 different parameters, the first is the location of the file we wish to write to, the second is our mydata object, and the third is the FileMode, which represents our file’s mode and permission bits.熱門LINE官方貼圖| LINE STORELINE STORE是LINE的官方線上商店。在此可購買熊大或兔兔等LINE卡通明星貼圖，或是其他熱門動漫貼圖。其中，動態貼圖或以流行搞笑梗令人發笑的有聲貼圖也 ...")
	err := ioutil.WriteFile("READ_FILE/localfile.txt", mydata, 0)
	if err != nil {
		fmt.Println(err)
	}

	data, errors := ioutil.ReadFile("READ_FILE/localfile.txt")
	if errors != nil {
		fmt.Println(errors)
	}
	fmt.Println(string(data))
}

func read() {
	data, err := ioutil.ReadFile("READ_FILE/localfile.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func open() {
	f, err := os.OpenFile("READ_FILE/localfile.txt", os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.WriteString("new data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\nnew data that wasn't there originally\n"); err != nil {
		panic(err)
	}

	data,err:=ioutil.ReadFile("READ_FILE/localfile.txt")
	if err !=nil{
		panic(err)
	}
	fmt.Println(string(data))
}
