package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello world", request.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
