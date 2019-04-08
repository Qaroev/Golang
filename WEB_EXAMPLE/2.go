package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello world", nil)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
