package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutext = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutext.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutext.Unlock()
}

func main() {
	http.HandleFunc("/", echoString)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "HI")
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
