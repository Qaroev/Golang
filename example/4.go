package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)


func main() {
	response, err := http.Get("https://raw.githubusercontent.com/Qaroev/App/master/data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}
